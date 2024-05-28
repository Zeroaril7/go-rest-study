package service

import (
	"context"

	"github.com/Zeroaril7/go-rest-study/model/web/request"
	"github.com/Zeroaril7/go-rest-study/model/web/response"
)

type CategoryService interface {
	Create(ctx context.Context, req request.CategoryCreateReq) response.CategoryResponse
	Update(ctx context.Context, req request.CategoryUpdateReq) response.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	GetById(ctx context.Context, categoryId int) response.CategoryResponse
	GetAll(ctx context.Context) []response.CategoryResponse
}
