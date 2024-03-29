// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
}

type Response struct {
	Message string `json:"message"`
	Data any `json:"data"`
}

type GetUserRequest struct {
	Id string `path:"id"`
}

type GetUserResponse struct {
	Message string `json:"message"`
	Data any `json:"data"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Data any `json:"data"`
}