package ocp_problem_api

import (
	"fmt"
	"github.com/ozoncp/ocp-problem-api/internal/repo"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
	"github.com/rs/zerolog"
	"os"
)

type OcpProblemAPI struct {
	desc.UnimplementedOcpProblemServer
	repo repo.RepoRemover
	log zerolog.Logger
}

func (pa *OcpProblemAPI) logResult(method string, request, result interface{}) {
	pa.log.Info().
		Str("method", method).
		Str("request", fmt.Sprintf("%v", request)).
		Msgf("%v", result)
}

func (pa *OcpProblemAPI) logError(method string, request interface{}, err error) {
	pa.log.Error().
		Str("method", method).
		Str("request", fmt.Sprintf("%v", request)).
		Msg(err.Error())
}

func NewOcpProblemAPI() desc.OcpProblemServer {
	return &OcpProblemAPI{
		repo: repo.NewRepo(),
		log: zerolog.New(os.Stdout),
	}
}