package controller

import (
	"net/http"
	"strconv"

	"github.com/Zeroaril7/go-rest-study/helper"
	req "github.com/Zeroaril7/go-rest-study/model/web/request"
	res "github.com/Zeroaril7/go-rest-study/model/web/response"
	"github.com/Zeroaril7/go-rest-study/service"
	"github.com/julienschmidt/httprouter"
)

type categoryControllerImpl struct {
	categoryService service.CategoryService
}

// Create implements CategoryController.
func (c *categoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	payload := new(req.CategoryCreateReq)

	helper.ReadFromReqBody(request, payload)

	result := c.categoryService.Create(request.Context(), *payload)

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResBody(writer, webResponse)

}

// Delete implements CategoryController.
func (c *categoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	c.categoryService.Delete(request.Context(), id)

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResBody(writer, webResponse)
}

// GetAll implements CategoryController.
func (c *categoryControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	result := c.categoryService.GetAll(request.Context())

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResBody(writer, webResponse)
}

// GetById implements CategoryController.
func (c *categoryControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	result := c.categoryService.GetById(request.Context(), id)

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResBody(writer, webResponse)
}

// Update implements CategoryController.
func (c *categoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	payload := new(req.CategoryUpdateReq)

	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)
	payload.Id = id

	helper.ReadFromReqBody(request, payload)

	result := c.categoryService.Update(request.Context(), *payload)

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResBody(writer, webResponse)
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryControllerImpl{categoryService: categoryService}
}
