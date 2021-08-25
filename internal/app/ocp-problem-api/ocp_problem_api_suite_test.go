package ocp_problem_api_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOcpProblemApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OcpProblemApi Suite")
}
