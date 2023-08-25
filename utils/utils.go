package utils

import "github.com/gin-gonic/gin"

type Handler interface {
	Handle() func(ctx *gin.Context)
}
