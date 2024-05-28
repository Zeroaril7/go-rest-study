package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Zeroaril7/go-rest-study/helper"
	"github.com/Zeroaril7/go-rest-study/model/domain"
)

type categoryRepositoryImpl struct {
}

// Add implements CategoryRepository.
func (r *categoryRepositoryImpl) Add(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "insert into category(name) values (?)"

	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

// Delete implements CategoryRepository.
func (r *categoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}

// GetAll implements CategoryRepository.
func (r *categoryRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) (result []domain.Category, err error) {
	sql := "select id, name from category"

	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)

	defer rows.Close()

	for rows.Next() {
		category := domain.Category{}
		err = rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		result = append(result, category)
	}

	return
}

// GetById implements CategoryRepository.
func (r *categoryRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id int) (result domain.Category, err error) {
	sql := "select id, name from category where id = ?"

	rows, err := tx.QueryContext(ctx, sql, id)
	helper.PanicIfError(err)

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&result.Id, &result.Name)
		helper.PanicIfError(err)
		if result == (domain.Category{}) {
			err = errors.New("Not Found")
		}
		return result, err
	}

	return
}

// Update implements CategoryRepository.
func (r *categoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) (result domain.Category, err error) {
	sql := "update category set name = ? where id = ?"

	_, err = tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category, err
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepositoryImpl{}
}
