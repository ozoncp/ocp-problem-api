package flusher_test

import (
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
	)

	Context("Flush", func() {
		It("", func() {
			solution := []utils.Problem{
				{},
				{},
				{},
				{},
				{},
				{},
				{},
			}

			ctrl = gomock.NewController(GinkgoT())
			mockRepo = mocks.NewMockRepo(ctrl)

			defer ctrl.Finish()

			problemFlusher := flusher.NewFlusher(2, mockRepo)

			mockRepo.EXPECT().AddEntities([]utils.Problem{{},{}}).Times(2).Return(nil)
			mockRepo.EXPECT().AddEntities([]utils.Problem{{},{}}).Times(3).Return(errors.New("some error"))
			mockRepo.EXPECT().AddEntities([]utils.Problem{{}}).Return(errors.New("some error"))

			result := problemFlusher.Flush([]utils.Problem{
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
