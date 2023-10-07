package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := struct {
		role string
	}{}

	ctx.BindJSON(&request)

	logger.Infof("request received: %+v", request)
}
