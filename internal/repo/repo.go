package repo

import (
	dbSql "database/sql"

	"github.com/Masterminds/squirrel"
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
	qb squirrel.StatementBuilderType
}

func NewObligationRepository(db *sqlx.DB) Repo {
	return &ObligationRepository{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *ObligationRepository) RemoveEntity(entityId uint64) error {
	sql, args, err := r.qb.
		Delete(table).
		Where(squirrel.Eq{"id": entityId}).
		ToSql()

	if err != nil {
		return err
	}

	result, err := r.db.Exec(sql, args...)
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
	sql, args, err := r.qb.Select("id, title, description").
		From(table).
		Limit(limit).
		Offset(offset).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	var obligations []entity.Obligation
	for rows.Next() {
		var obligation entity.Obligation
		err := rows.Scan(&obligation.ID, &obligation.Title, &obligation.Description)
		if err != nil {
			return nil, err
		}

		obligations = append(obligations, obligation)
	}

	return obligations, nil
}

func (r *ObligationRepository) AddEntity(entity *entity.Obligation) error {
	sql, args, err := r.qb.
		Insert(table).Columns("title", "description").
		Values(entity.Title, entity.Description).
		ToSql()

	if err != nil {
		return err
	}

	var id int
	err = r.db.QueryRow(sql+" RETURNING id", args...).Scan(&id)
	if err != nil {
		return err
	}

	entity.ID = uint(id)

	return nil
}

func (r ObligationRepository) AddEntities(entities []*entity.Obligation) error {
	t, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, entity := range entities {
		if err := r.AddEntity(entity); err != nil {
			t.Rollback()
			return err
		}
	}

	return t.Commit()
}

func (r *ObligationRepository) DescribeEntity(entityId uint64) (*entity.Obligation, error) {
	sql, args, err := r.qb.Select("id, title, description").
		From(table).
		Where(squirrel.Eq{"id": entityId}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var obligation entity.Obligation

	row := r.db.QueryRow(sql, args...)
	err = row.Scan(&obligation.ID, &obligation.Title, &obligation.Description)
	if err != nil {
		return nil, err
	}

	return &obligation, nil
}
