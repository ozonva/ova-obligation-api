package flusher

import (
	"github.com/ozonva/ova-obligation-api/internal/entity"
	"github.com/ozonva/ova-obligation-api/internal/repo"
)

type Flusher interface {
	Flush(entities []entity.Obligation) []entity.Obligation
}

func NewFlusher(repository repo.Repo) Flusher {
	return &flusher{repository: repository}
}

type flusher struct {
	repository repo.Repo
}

func (f *flusher) Flush(entities []entity.Obligation) []entity.Obligation {
	var failed []entity.Obligation
	for _, entity := range entities {
		if err := f.repository.AddEntity(entity); err != nil {
			failed = append(failed, entity)
		}
	}

	return failed
}
