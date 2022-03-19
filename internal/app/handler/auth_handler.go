package handler

import "github.com/gin-gonic/gin"

type AuthHandler struct {

}

func ProvideAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (a AuthHandler) Signup(c *gin.Context) {

}

func (a AuthHandler) Login(c *gin.Context) {

}