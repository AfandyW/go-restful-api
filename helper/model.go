package helper

import (
	"go-restful-api/model/domain"
	"go-restful-api/model/web"
	"net/http"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoriesResponse(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ResponseCategoryNotNil(category web.CategoryResponse) (webResponse web.WebResponse) {
	if category.Id == 0 {
		webResponse = web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		}
	} else {
		webResponse = web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   category,
		}
	}
	return
}
