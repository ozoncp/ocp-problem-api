package main

import (
	"flag"
	ocp_problem_api "github.com/ozoncp/ocp-problem-api/internal/app/ocp-problem-api"
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

	serviceRunner := ocp_problem_api.NewRunner(grpcPort, restPort, *serverHost, ocp_problem_api.NewOcpProblemAPI())
	if err := serviceRunner.Run(); err != nil {

	}
}