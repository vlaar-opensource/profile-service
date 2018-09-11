package app

import (
	"net/http"

	"github.com/vlaar-opensource/profile-service/helper"
	"github.com/vlaar-opensource/profile-service/messages"
)

//Register blablabla
func (app *App) Register(w http.ResponseWriter, r *http.Request) {

	req := &messages.RegisterRequest{}
	err := helper.GetRequest(r, req)
	if err != nil {
		helper.WriteResponse(w, messages.ErrorResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Invalid Parameter",
		})
		return
	}

	// do the processing..

	helper.WriteResponse(w, messages.OkResponse{
		Message: "OK",
	})

}

//CheckUsername blablabla
func (app *App) CheckUsername(w http.ResponseWriter, r *http.Request) {

}

//Login blablabal
func (app *App) Login(w http.ResponseWriter, r *http.Request) {

}
