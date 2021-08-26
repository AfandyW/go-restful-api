package repository

import (
	"context"
	"database/sql"
	"go-restful-api/helper"
	"go-restful-api/model/domain"
)

type categoryRepository struct {
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (c *categoryRepository) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "INSERT INTO category(name) values($1) RETURNING id"
	err := tx.QueryRowContext(ctx, sql, category.Name).Scan(&category.Id)
	helper.PanicIfError(err)

	return category
}

func (c *categoryRepository) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "update category set name = $1 where id = $2"

	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}
func (c *categoryRepository) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "delete from category where id = $1"

	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}
func (c *categoryRepository) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "SELECT id, name from category WHERE id=$1"
	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)

	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(
			&category.Id,
			&category.Name)

		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, nil
	}
}
func (c *categoryRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error) {
	sql := "SELECT id, name FROM category"

	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)

	defer rows.Close()
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
