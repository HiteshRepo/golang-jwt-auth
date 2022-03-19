package handler

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func ProvideUserHandler() UserHandler {
	return UserHandler{}
}

func (u UserHandler) ListUsers(c *gin.Context) {
	c.JSON(200, gin.H{"success": "Listing users..."})
}
