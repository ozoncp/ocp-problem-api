package main

import (
	"flag"
	"fmt"
	ocp "github.com/ozoncp/ocp-problem-api/internal/app/ocp-problem-api"
	"github.com/ozoncp/ocp-problem-api/internal/repo"
	"github.com/rs/zerolog"
	"io"
	"os"
	"os/signal"
	"syscall"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

var (
	serverHost = flag.String("host",  "", "server host endpoint")
	serverGrpcPort = flag.Int("grpc-port", 8082, "server GRPC port")
	serverRestPort = flag.Int("rest-port", 8083, "server REST API port")
	serverMetricPort = flag.Int("metric-port", 9100, "metric port")
)

func main()  {
	flag.Parse()
	grpcPort := uint32(*serverGrpcPort)
	restPort := uint32(*serverRestPort)
	metricPort := uint32(*serverMetricPort)

	dbRepo := repo.NewPgRepo(os.Getenv("DATABASE_URL"))
	kafkaRepo := repo.NewRepoKafka([]string{os.Getenv("KAFKA_BROKER")})
	serviceRepo := repo.NewRepoChain(dbRepo, kafkaRepo)

	logger := zerolog.New(os.Stdout)
	tracerCloser, err := initTracer()
	defer tracerCloser.Close()

	if err != nil {
		logger.Error().Msg(err.Error())
	}

	service := ocp.NewOcpProblemAPI(serviceRepo, logger, 2)
	serviceRunner := ocp.NewRunner(
		grpcPort,
		restPort,
		metricPort,
		*serverHost,
		service,
		logger,
		)

	defer serviceRunner.Stop()

	startSignalCatcher(serviceRunner, tracerCloser)
	if err := serviceRunner.Run(); err != nil {
		fmt.Println(err.Error())
	}
}

func initTracer() (closer io.Closer, err error) {
	cfg := jaegercfg.Configuration{
		ServiceName: "ocp-problem-api",
		Sampler:     &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter:    &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	if err != nil {
		return
	}

	opentracing.SetGlobalTracer(tracer)
	return
}

func startSignalCatcher(runner ocp.PublicRunner, tracerCloser io.Closer) {
	go func() {
		defer runner.Stop()
		defer tracerCloser.Close()

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
	}()
}