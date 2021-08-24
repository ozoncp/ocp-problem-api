package main

import (
	"flag"
	"fmt"
	ocp "github.com/ozoncp/ocp-problem-api/internal/app/ocp-problem-api"
	"github.com/ozoncp/ocp-problem-api/internal/repo"
	"github.com/rs/zerolog"
	"os"
)

var (
	serverHost = flag.String("host",  "", "server host endpoint")
	serverGrpcPort = flag.Int("grpc-port", 8082, "server GRPC port")
	serverRestPort = flag.Int("rest-port", 8083, "server REST API port")
)

func main()  {
	flag.Parse()
	grpcPort := uint32(*serverGrpcPort)
	restPort := uint32(*serverRestPort)

	serviceRepo, err := repo.NewPgRepo(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//serviceRepo := repo.NewFakeRepo()
	logger := zerolog.New(os.Stdout)
	service := ocp.NewOcpProblemAPI(serviceRepo, logger)
	serviceRunner := ocp.NewRunner(
		grpcPort,
		restPort,
		*serverHost,
		service,
		logger,
		)
	if err := serviceRunner.Run(); err != nil {
		fmt.Println(err.Error())
	}
}