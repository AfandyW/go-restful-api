package controller

import (
	"encoding/json"
	"go-restful-api/helper"
	"go-restful-api/model/web"
	"go-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type categoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(cs service.CategoryService) CategoryController {
	return &categoryControllerImpl{
		CategoryService: cs,
	}
}

func (controller *categoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "applications/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *categoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "Accepted",
		Data:   categoryResponse,
	}

	helper.WriterToResponseBody(writer, webResponse)
}
func (controller *categoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "Accepted",
		Data:   nil,
	}

	helper.WriterToResponseBody(writer, webResponse)
}
func (controller *categoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	webResponse := helper.ResponseCategoryNotNil(categoryResponse)

	helper.WriterToResponseBody(writer, webResponse)
}
func (controller *categoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponse := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoriesResponse,
	}

	helper.WriterToResponseBody(writer, webResponse)
}
