package service

import (
	"context"
	"database/sql"
	"go-restful-api/helper"
	"go-restful-api/model/domain"
	"go-restful-api/model/web"
	"go-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type categoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepo repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &categoryServiceImpl{
		CategoryRepository: categoryRepo,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *categoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback(tx)
	}()

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *categoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback(tx)
	}()

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}
func (service *categoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback(tx)
	}()

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}
func (service *categoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback(tx)
	}()
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}
func (service *categoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback(tx)
	}()

	categories, err := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoriesResponse(categories)
}