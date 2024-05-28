package repository

import (
	"context"
	"database/sql"

	"github.com/Zeroaril7/go-rest-study/model/domain"
)

type CategoryRepository interface {
	GetById(ctx context.Context, tx *sql.Tx, id int) (result domain.Category, err error)
	GetAll(ctx context.Context, tx *sql.Tx) (result []domain.Category, err error)
	Add(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) (result domain.Category, err error)
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
}
