package ocp_problem_api

import (
	"fmt"
	"github.com/ozoncp/ocp-problem-api/internal/flusher"
	"github.com/ozoncp/ocp-problem-api/internal/repo"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
	"github.com/rs/zerolog"
)

type OcpProblemAPI struct {
	desc.UnimplementedOcpProblemServer
	repo repo.RepoRemover
	log zerolog.Logger
	flusher flusher.Flusher
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

func NewOcpProblemAPI(r repo.RepoRemover, logger zerolog.Logger, chunkSize int) desc.OcpProblemServer {
	return &OcpProblemAPI{
		repo: r,
		log: logger,
		flusher: flusher.NewFlusher(chunkSize, r),
	}
}