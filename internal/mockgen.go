package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-problem-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/repo_remover_mock.go -package=mocks github.com/ozoncp/ocp-problem-api/internal/repo RepoRemover
