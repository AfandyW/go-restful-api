package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restful-api/helper"
	"go-restful-api/model/domain"
)

type categoryRepository struct {
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (c *categoryRepository) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "insert into category(name) value(?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)
	category.Id = int(id)

	return category
}

func (c *categoryRepository) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "update category set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}
func (c *categoryRepository) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "update from category where id = ?"

	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}
func (c *categoryRepository) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "select id, name from category where id = ?"

	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(
			&category.Id,
			&category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Category Not Found")
	}
}
func (c *categoryRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error) {
	sql := "select id, name from category"

	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)

	categories := []domain.Category{}
	for rows.Next() {
		category := domain.Category{}

		err := rows.Scan(
			&category.Id,
			&category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories, nil
}
