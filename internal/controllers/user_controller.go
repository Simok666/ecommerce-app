package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"message": "You are authenticated",
	})
}
