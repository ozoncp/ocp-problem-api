package flusher_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-problem-api/internal/flusher"
	"github.com/ozoncp/ocp-problem-api/internal/mocks"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
)

var _ = Describe("Flusher", func() {

	var (
		ctrl *gomock.Controller
		mockRepo *mocks.MockRepo
		result []utils.Problem
		ctx context.Context
	)

	ctx = context.Background()

	Context("Flush", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockRepo = mocks.NewMockRepo(ctrl)
		})

		AfterEach(func() {
			ctrl.Finish()
		})

		It("Check nil problem list", func() {
			problemFlusher := flusher.NewFlusher(2, mockRepo)
			result = problemFlusher.Flush(ctx,nil)

			Expect(result).To(BeNil())
		})

		It("Check empty problem list", func() {
			problemFlusher := flusher.NewFlusher(2, mockRepo)
			result = problemFlusher.Flush(ctx, []utils.Problem{})

			Expect(result).To(BeNil())
		})

		It("Check bulk save entities", func() {
			solution := []utils.Problem{
				{},
				{},
				{},
				{},
				{},
				{},
				{},
			}

			problemFlusher := flusher.NewFlusher(2, mockRepo)
			mockRepo.EXPECT().AddEntities(ctx, []utils.Problem{{},{}}).Times(2).Return(nil)
			mockRepo.EXPECT().AddEntities(ctx, []utils.Problem{{},{}}).Times(3).Return(errors.New("some error"))
			mockRepo.EXPECT().AddEntities(ctx, []utils.Problem{{}}).Return(errors.New("some error"))

			result = problemFlusher.Flush(ctx, []utils.Problem{
				{},
				{},
				{},
				{},
				{},
				{},
				{},
				{},
				{},
				{},
				{},
			})

			Expect(result).To(Equal(solution))
		})
	})
})
