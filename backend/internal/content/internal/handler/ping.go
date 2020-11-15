package handler

import "github.com/gin-gonic/gin"

func Ping(_ *gin.Context) (resp interface{}, err error) {
	return "pong", nil
}
