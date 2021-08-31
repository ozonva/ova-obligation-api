package repo

import (
	dbSql "database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-obligation-api/internal/entity"
)

const table = "obligation"

type Repo interface {
	AddEntities(entities []*entity.Obligation) error
	AddEntity(entity *entity.Obligation) error
	ListEntities(limit, offset uint64) ([]entity.Obligation, error)
	DescribeEntity(entityId uint64) (*entity.Obligation, error)
	RemoveEntity(entityId uint64) error
}

type ObligationRepository struct {
	db *sqlx.DB
}

func NewObligationRepository(db *sqlx.DB) Repo {
	return &ObligationRepository{
		db: db,
	}
}

func (r *ObligationRepository) RemoveEntity(entityId uint64) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table)
	result, err := r.db.Exec(sql, entityId)
	if err != nil {
		return err
	}

	rowEffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowEffected == 0 {
		return dbSql.ErrNoRows
	}

	return nil
}

func (r *ObligationRepository) ListEntities(limit, offset uint64) ([]entity.Obligation, error) {
	sql := fmt.Sprintf("SELECT * FROM %s LIMIT $1 OFFSET $2", table)
	var obligations []entity.Obligation
	err := r.db.Select(&obligations, sql, limit, offset)
	if err != nil {
		return nil, err
	}

	return obligations, nil
}

func (r *ObligationRepository) AddEntity(entity *entity.Obligation) error {
	err := r.saveEntity(r.db, entity)
	if err != nil {
		return err
	}

	return nil
}

func (r *ObligationRepository) saveEntity(query sqlx.Queryer, entity *entity.Obligation) error {
	var id int
	sql := fmt.Sprintf(`INSERT INTO %s (title,description) VALUES ($1,$2) RETURNING id`, table)
	err := query.QueryRowx(sql, entity.Title, entity.Description).Scan(&id)
	if err != nil {
		return err
	}

	entity.ID = uint(id)

	return nil
}

func (r *ObligationRepository) AddEntities(entities []*entity.Obligation) error {
	tx := r.db.MustBegin()
	for _, entity := range entities {
		if err := r.saveEntity(tx, entity); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *ObligationRepository) DescribeEntity(entityId uint64) (*entity.Obligation, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table)
	var obligation entity.Obligation
	err := r.db.Get(&obligation, sql, entityId)
	if err != nil {
		return nil, err
	}

	return &obligation, nil
}
