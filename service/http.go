/*
HTTP接口服务

BetaX Page To What
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/

package service

import (
	"p2w/global"

	"github.com/gin-gonic/gin"
)

type HTTP struct{}

func (http HTTP) ApiToPDF(ctx *gin.Context) {
	// 地址
	url := ctx.Query("url")
	if len(url) == 0 {
		global.ReturnMessage(ctx, false, "url cannot be empty")
		return
	}
	// 访问地址截取页面内容
	cache := ToPDF(url)
	// 发送接口(填了就不放回PDF而是直接发送)
	send := ctx.Query("send")
	if len(url) == 0 {
		global.ReturnFile(ctx, "application/pdf", cache)
	} else {
		// 标识码
		code := global.GetCode(ctx.Query("code"))
		global.SendFile(send, cache, code+".pdf")
		global.ReturnSuccess(ctx, true)
	}
}

func (http HTTP) ApiToImage(ctx *gin.Context) {
	// 地址
	url := ctx.Query("url")
	if len(url) == 0 {
		global.ReturnMessage(ctx, false, "url cannot be empty")
		return
	}
	// 截取元素
	element := ctx.Query("element")
	// 图片质量
	quality := ctx.Query("quality")
	var cache []byte
	if len(url) == 0 {
		cache = ToFullImage(url, global.ToInt(quality, 90))
	} else {
		cache = ToElementImage(url, element)
	}
	// 发送接口(填了就不放回PDF而是直接发送)
	send := ctx.Query("send")
	if len(url) == 0 {
		global.ReturnFile(ctx, "image/png", cache)
	} else {
		// 标识码
		code := global.GetCode(ctx.Query("code"))
		global.SendFile(send, cache, code+".png")
		global.ReturnSuccess(ctx, true)
	}
}
