/*
全局公共响应

BetaX Page To What
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/

package global

import (
	"time"

	"github.com/gin-gonic/gin"
)

func ReturnError(ctx *gin.Context, err CustomError) {
	ctx.JSON(200, err)
	ctx.Abort()
}

type commonResponse struct {
	State   bool   `json:"bool"`
	Message string `json:"message"`
	Time    int64  `json:"time"`
}

func ReturnMessage(ctx *gin.Context, state bool, message string) {
	ctx.JSON(200, commonResponse{
		State:   state,
		Message: message,
		Time:    time.Now().Unix(),
	})
	ctx.Abort()
}

func ReturnSuccess(ctx *gin.Context, obj any) {
	ctx.JSON(200, obj)
	ctx.Abort()
}

func ReturnFile(ctx *gin.Context, type_ string, data []byte) {
	ctx.Header("Content-Type", type_)
	ctx.Writer.Write(data)
	ctx.Abort()
}
