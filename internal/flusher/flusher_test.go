package flusher

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/onsi/gomega"
	"github.com/ozonva/ova-obligation-api/internal/entity"
	"github.com/ozonva/ova-obligation-api/internal/repo"
)

func Test_flusher_Flush(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := repo.NewMockRepo(ctrl)
	m.EXPECT().AddEntity(entity.Obligation{Title: "test", Description: "test", ID: 1}).Return(nil)

	flusher := New(m)

	g := gomega.NewGomegaWithT(t)

	result := flusher.Flush([]entity.Obligation{{Title: "test", Description: "test", ID: 1}})
	g.Ω(result).Should(gomega.BeNil())
}

func Test_flusher_FlushWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := repo.NewMockRepo(ctrl)
	m.EXPECT().AddEntity(entity.Obligation{Title: "test", Description: "test", ID: 1}).Return(errors.New("test"))

	flusher := New(m)

	g := gomega.NewGomegaWithT(t)

	result := flusher.Flush([]entity.Obligation{{Title: "test", Description: "test", ID: 1}})
	expected := []entity.Obligation{{Title: "test", Description: "test", ID: 1}}
	g.Ω(result).Should(gomega.Equal(expected))
}
