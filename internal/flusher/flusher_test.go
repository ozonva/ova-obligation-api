package flusher

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-obligation-api/internal/entity"
	"github.com/ozonva/ova-obligation-api/internal/repo"
)

var _ = Describe("Flusher", func() {
	var (
		mockCtrl    *gomock.Controller
		mockRepo    *repo.MockRepo
		testFlusher Flusher
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockRepo(mockCtrl)
		testFlusher = NewFlusher(mockRepo)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Writing data to storage", func() {
		When("Write data", func() {
			Context("Write without error", func() {
				BeforeEach(func() {
					mockRepo.EXPECT().AddEntity(entity.Obligation{ID: 1}).Return(nil)
				})
				It("Should return nil", func() {
					result := testFlusher.Flush([]entity.Obligation{{ID: 1}})
					Expect(result).Should(BeNil())
				})
			})
			Context("Write with error", func() {
				BeforeEach(func() {
					mockRepo.EXPECT().AddEntity(entity.Obligation{ID: 1}).Return(errors.New("test"))
				})
				It("Should return not saved entity", func() {
					result := testFlusher.Flush([]entity.Obligation{{ID: 1}})
					expected := []entity.Obligation{{ID: 1}}
					Expect(result).Should(Equal(expected))
				})
			})
		})
	})
})
