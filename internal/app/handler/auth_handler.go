package handler

import "github.com/gin-gonic/gin"

type AuthHandler struct {

}

func ProvideAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (a AuthHandler) FetchToken(c *gin.Context) {
	c.JSON(200, gin.H{"success": "Fetching token..."})
}