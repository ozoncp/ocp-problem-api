package ocp_problem_api_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	ocp_problem_api "github.com/ozoncp/ocp-problem-api/internal/app/ocp-problem-api"
	"github.com/ozoncp/ocp-problem-api/internal/mocks"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
	desc "github.com/ozoncp/ocp-problem-api/pkg/ocp-problem-api"
	"github.com/rs/zerolog"
	"io/ioutil"
)

var _ = Describe("Server", func() {
	ctx := context.Background()
	ctrl := gomock.NewController(GinkgoT())
	repoRemover := mocks.NewMockRepoRemover(ctrl)
	logger := zerolog.New(ioutil.Discard)
	service := ocp_problem_api.NewOcpProblemAPI(repoRemover, logger)

	Context("CreateProblemV1", func() {
		problemV1 := &desc.ProblemV1{
			Id: 100,
			UserId: 2,
			Text: "some text",
		}

		problemList := []utils.Problem{
			{Id: problemV1.Id, UserId: problemV1.UserId, Text: problemV1.Text},
		}

		expectResult1 := desc.ResultSaveV1{
			Id: problemV1.Id,
		}

		It("Success create", func() {
			repoRemover.EXPECT().AddEntities(ctx, problemList).Times(1).Return(nil)

			result, err := service.CreateProblemV1(ctx, problemV1)
			gomega.Expect(result.Id).To(gomega.Equal(expectResult1.Id))
			gomega.Expect(err).To(gomega.BeNil())

		})

		It("Fail create", func() {
			repoRemover.EXPECT().AddEntities(ctx, problemList).Times(1).Return(errors.New("duplicate problem"))
			result, err := service.CreateProblemV1(ctx, problemV1)
			gomega.Expect(result).To(gomega.BeNil())
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
	})

	Context("DescribeProblemV1", func() {
		It("Success describe", func() {
			problemId := uint64(2)
			problem := &utils.Problem{Id: problemId, UserId: 2, Text: "someText"}
			problemQueryV1 := &desc.ProblemQueryV1{Id: problemId}

			repoRemover.EXPECT().DescribeEntity(ctx, problemId).Times(1).Return(problem, nil)

			problemV1, err := service.DescribeProblemV1(ctx, problemQueryV1)
			gomega.Expect(problemV1.Id).To(gomega.Equal(problem.Id))
			gomega.Expect(problemV1.UserId).To(gomega.Equal(problem.UserId))
			gomega.Expect(problemV1.Text).To(gomega.Equal(problem.Text))
			gomega.Expect(err).To(gomega.BeNil())
		})

		It("Fail describe", func() {
			problemId := uint64(2222)
			problemQueryV1 := &desc.ProblemQueryV1{Id: problemId}
			errorText := "some error"

			repoRemover.
				EXPECT().
				DescribeEntity(ctx, problemId).
				Times(1).
				Return(nil, errors.New(errorText))

			problemV1, err := service.DescribeProblemV1(ctx, problemQueryV1)
			gomega.Expect(problemV1).To(gomega.BeNil())
			gomega.Expect(err).NotTo(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal(errorText))
		})
	})

	Context("ListProblemsV1", func() {
		It("Not empty list", func() {
			limit := uint64(10)
			offset := uint64(20)
			problemList := []utils.Problem{
				{Id: 33, UserId: 2, Text: "some text"},
			}
			problemListQueryV1 := &desc.ProblemListQueryV1{Limit: limit, Offset: offset}

			repoRemover.EXPECT().ListEntities(ctx, limit, offset).Return(problemList, nil)

			resultList, err :=  service.ListProblemsV1(ctx, problemListQueryV1)
			gomega.Expect(len(resultList.List)).To(gomega.Equal(1))
			gomega.Expect(resultList.List[0].Id).To(gomega.Equal(problemList[0].Id))
			gomega.Expect(err).To(gomega.BeNil())
		})

		It("Empty list", func() {
			limit := uint64(10)
			offset := uint64(2000)
			textError := "one more error"
			problemListQueryV1 := &desc.ProblemListQueryV1{Limit: limit, Offset: offset}
			repoRemover.EXPECT().ListEntities(ctx, limit, offset).Return(nil, errors.New(textError))

			resultList, err := service.ListProblemsV1(ctx, problemListQueryV1)
			gomega.Expect(resultList).To(gomega.BeNil())
			gomega.Expect(err).NotTo(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal(textError))
		})
	})

	Context("RemoveProblemV1", func() {
		It("Success remove", func() {
			problemId := uint64(5)
			problemQueryV1 := &desc.ProblemQueryV1{Id: problemId}

			repoRemover.EXPECT().RemoveEntity(ctx, problemId).Return(nil)

			result, err :=  service.RemoveProblemV1(ctx, problemQueryV1)
			gomega.Expect(result.Id).To(gomega.Equal(problemId))
			gomega.Expect(result.IsSuccess).To(gomega.Equal(true))
			gomega.Expect(err).To(gomega.BeNil())
		})
		It("Fail remove", func() {
			problemId := uint64(77)
			problemQueryV1 := &desc.ProblemQueryV1{Id: problemId}
			errorText := "problem not found"

			repoRemover.EXPECT().RemoveEntity(ctx, problemId).Return(errors.New(errorText))

			result, err :=  service.RemoveProblemV1(ctx, problemQueryV1)
			gomega.Expect(result).To(gomega.BeNil())
			gomega.Expect(err).NotTo(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal(errorText))
		})
	})
})
