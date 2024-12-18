package models

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Message interface{} `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, Response{message})
}
