package ocp_problem_api

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type Stopper interface {
	stop() error
}

type GrpcRunner interface {
	Stopper
	runGrpc(port uint32, host string, service interface{}) error
}

type RestRunner interface {
	Stopper
	runRest(port, grpcPort uint32, host string) error
}

type PublicRunner interface {
	SetRestRunner(runner RestRunner) error
	SetGrpcRunner(runner GrpcRunner) error
	Run() error
	Stop() error
}

type defaultGrpcRunner struct {
	server *grpc.Server
}

func (pgr *defaultGrpcRunner) runGrpc(port uint32, host string, service interface{}) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%v:%d", host, port))
	log.Printf("Listening GRPC on %v:%d", host, port)

	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
		return err
	}

	ocpService, ok := service.(desc.OcpProblemServer)
	if !ok {
		return errors.New("invalid service type")
	}

	if pgr.server == nil {
		pgr.server = grpc.NewServer()
	}

	desc.RegisterOcpProblemServer(pgr.server, ocpService)
	if err := pgr.server.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
		return err
	}

	return nil
}

func (pgr *defaultGrpcRunner) stop() error {
	if pgr.server == nil {
		return errors.New("server is not running")
	}

	pgr.server.Stop()
	return nil
}

type defaultRestRunner struct {}

func (prr *defaultRestRunner) runRest(port, grpcPort uint32, host string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpProblemHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%v:%d", host, grpcPort), opts)
	if err != nil {
		return err
	}

	log.Printf("Listening REST on %v:%d", host, port)

	return http.ListenAndServe(fmt.Sprintf("%v:%d", host, port), mux)
}

func (prr *defaultRestRunner) stop() error {
	return nil
}

type projectRunner struct {
	rest RestRunner
	grpc GrpcRunner
	grpcPort uint32
	restPort uint32
	host string
	isRunning bool
	service interface{}
	log zerolog.Logger
}

func (pr *projectRunner) Stop() error {
	var stopError error

	if err := pr.rest.stop(); err != nil {
		stopError = utils.NewWrappedError(err.Error(), stopError)
		pr.log.Error().Msg(err.Error())
	}

	if err := pr.grpc.stop(); err != nil {
		stopError = utils.NewWrappedError(err.Error(), stopError)
		pr.log.Error().Msg(err.Error())
	}

	return stopError
}

func (pr *projectRunner) SetRestRunner(runner RestRunner) error {
	if pr.isRunning {
		return errors.New("service is already running")
	}

	pr.rest = runner
	return nil
}

func (pr *projectRunner) SetGrpcRunner(runner GrpcRunner) error {
	if pr.isRunning {
		return errors.New("service is already running")
	}

	pr.grpc = runner
	return nil
}

func (pr *projectRunner) Run() error {
	errChan := make(chan error)
	defer close(errChan)

	pr.isRunning = false
	go func(errChan chan <-error) {
		if err := pr.grpc.runGrpc(pr.grpcPort, pr.host, pr.service); err != nil {
			errChan <-err
		}
	}(errChan)


	go func(errChan chan <-error) {
		if err := pr.rest.runRest(pr.restPort, pr.grpcPort, pr.host); err != nil {
			if err := pr.grpc.stop(); err != nil {
				pr.log.Error().Msg(err.Error())
			}

			errChan <-err
		}
	}(errChan)

	pr.isRunning = true

	err := <- errChan
	pr.log.Error().Msg(err.Error())

	return err
}

func NewRunner(grpcPort uint32, restPort uint32, host string, service interface{}, logger zerolog.Logger) PublicRunner {
	return &projectRunner{
		rest: &defaultRestRunner{},
		grpc: &defaultGrpcRunner{},
		grpcPort: grpcPort,
		restPort: restPort,
		host: host,
		service: service,
		log: logger,
	}
}