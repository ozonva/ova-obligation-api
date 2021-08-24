package repo

import "github.com/ozonva/ova-obligation-api/internal/entity"

type Repo interface {
	AddEntities(entities []entity.Obligation) error
	AddEntity(entity entity.Obligation) error
	ListEntities(limit, offset uint64) ([]entity.Obligation, error)
	DescribeEntity(entityId uint64) (*entity.Obligation, error)
}
