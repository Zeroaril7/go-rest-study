package exception

import (
	"net/http"

	"github.com/Zeroaril7/go-rest-study/helper"
	"github.com/Zeroaril7/go-rest-study/model/web/response"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if NotFoundError(w, r, err) {
		return
	}

	InternalServerError(w, r, err)

}

func NotFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {

	exception, ok := err.(BaseError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := response.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResBody(w, res)
		return true
	} else {
		return false
	}

}

func InternalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	res := response.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResBody(w, res)
}
