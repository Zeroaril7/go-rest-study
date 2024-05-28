package service

import (
	"context"
	"database/sql"

	"github.com/Zeroaril7/go-rest-study/helper"
	"github.com/Zeroaril7/go-rest-study/model/domain"
	"github.com/Zeroaril7/go-rest-study/model/web/request"
	"github.com/Zeroaril7/go-rest-study/model/web/response"
	"github.com/Zeroaril7/go-rest-study/pkg/exception"
	"github.com/Zeroaril7/go-rest-study/repository"
	"github.com/go-playground/validator/v10"
)

type categoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

// Create implements CategoryService.
func (s *categoryServiceImpl) Create(ctx context.Context, req request.CategoryCreateReq) response.CategoryResponse {
	err := s.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx)

	category := domain.Category{
		Name: req.Name,
	}

	category = s.categoryRepository.Add(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

// Delete implements CategoryService.
func (s *categoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx)

	category, err := s.categoryRepository.GetById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	s.categoryRepository.Delete(ctx, tx, category)

}

// GetAll implements CategoryService.
func (s *categoryServiceImpl) GetAll(ctx context.Context) []response.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx)

	categories, err := s.categoryRepository.GetAll(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var response []response.CategoryResponse

	for _, category := range categories {
		response = append(response, helper.ToCategoryResponse(category))
	}

	return response
}

// GetById implements CategoryService.
func (s *categoryServiceImpl) GetById(ctx context.Context, categoryId int) response.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx)

	category, err := s.categoryRepository.GetById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

// Update implements CategoryService.
func (s *categoryServiceImpl) Update(ctx context.Context, req request.CategoryUpdateReq) response.CategoryResponse {
	err := s.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx)

	category := domain.Category{
		Id:   req.Id,
		Name: req.Name,
	}

	category, err = s.categoryRepository.Update(ctx, tx, category)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &categoryServiceImpl{categoryRepository: categoryRepository, DB: db, Validate: validate}
}
