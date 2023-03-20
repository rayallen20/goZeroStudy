package middlewares

import (
	"fmt"
	"net/http"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (u *UserMiddleware) LoginAndRegister(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("调用login和register之前执行...\n")
		next(w, r)
		fmt.Printf("调用login和register之后执行...\n")
	}
}

func (u *UserMiddleware) Global(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("调用所有函数之前执行...\n")
		next(w, r)
		fmt.Printf("调用所有函数之后执行...\n")
	}
}
