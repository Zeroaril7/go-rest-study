package helper

import (
	"github.com/Zeroaril7/go-rest-study/model/domain"
	"github.com/Zeroaril7/go-rest-study/model/web/response"
)

func ToCategoryResponse(category domain.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
