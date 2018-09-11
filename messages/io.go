package messages

import (
	"encoding/json"
)

//OkResponse ...
type OkResponse struct {
	Code    int         `json:"code:omitempty"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

// UnmarshalOkResponse blablabla
func UnmarshalOkResponse(b []byte) (OkResponse, error) {
	o := OkResponse{}
	err := json.Unmarshal(b, &o)
	return o, err
}

//Marshal blabla
func (r *OkResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//ErrorResponse blabla
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//UnmarshalErrorResponse blabla
func UnmarshalErrorResponse(b []byte) (ErrorResponse, error) {
	o := ErrorResponse{}
	err := json.Unmarshal(b, &o)
	return o, err
}

//Marshal blabla
func (r *ErrorResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//RegisterRequest blabla
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//UnmarshalRegisterRequest blabla
func UnmarshalRegisterRequest(b []byte) (RegisterRequest, error) {
	o := RegisterRequest{}
	err := json.Unmarshal(b, &o)
	return o, err
}

//Marshal blabla
func (r *RegisterRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//UsernameCheckRequest blabla
type UsernameCheckRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//UnmarshalUsernameCheckRequest blabla
func UnmarshalUsernameCheckRequest(b []byte) (UsernameCheckRequest, error) {
	o := UsernameCheckRequest{}
	err := json.Unmarshal(b, &o)
	return o, err
}

//Marshal blabla..
func (r *UsernameCheckRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//LoginRequest blabla
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//UnmarshalLoginRequest blabla
func UnmarshalLoginRequest(b []byte) (LoginRequest, error) {
	o := LoginRequest{}
	err := json.Unmarshal(b, &o)
	return o, err
}

//Marshal blabla
func (r *LoginRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
