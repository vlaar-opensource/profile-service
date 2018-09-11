package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vlaar-opensource/profile-service/helper"
	"github.com/vlaar-opensource/profile-service/messages"
	"golang.org/x/crypto/bcrypt"
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

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		helper.WriteResponse(w, messages.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Crypto Failed",
		})
	}
	// do the processing..
	obj, err := app.DB.AddUser(req.Username, string(pass))

	if err != nil {
		helper.WriteResponse(w, messages.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	obj.Password = "[REDACTED]"

	helper.WriteResponse(w, messages.OkResponse{
		Message: "OK",
		Payload: obj,
	})

}

//CheckUsername blablabla
func (app *App) CheckUsername(w http.ResponseWriter, r *http.Request) {

}

//Login blablabal
func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	req := messages.LoginRequest{}
	helper.GetRequest(r, &req)

	user, err := app.DB.GetByUsername(req.Username)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		helper.WriteResponse(w, messages.ErrorResponse{
			Code:    http.StatusForbidden,
			Message: err.Error(),
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uname": user.Username,
		"iss":   "vlaar",
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenStr, err := token.SignedString([]byte("supersecret"))
	if err != nil {
		helper.WriteResponse(w, messages.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprint("JWT Signing: ", err.Error()),
		})
		return
	}

	helper.WriteResponse(w, messages.OkResponse{
		Message: "OK",
		Payload: map[string]string{
			"jwtToken": tokenStr,
		},
	})
}
