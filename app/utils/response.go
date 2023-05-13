package utils

import (
	"github.com/goravel/framework/contracts/http"
)

func Success(ctx http.Context, data interface{}, msg string) {
	ctx.Response().Json(http.StatusOK, http.Json{
		"code": 200,
		"data": data,
		"msg":  msg,
	})
}

func Fail(ctx http.Context, data interface{}, msg string) {
	ctx.Response().Json(http.StatusOK, http.Json{
		"code": 400,
		"data": data,
		"msg":  msg,
	})
}

func Grant(ctx http.Context, data interface{}, msg string) {
	ctx.Response().Json(http.StatusOK, http.Json{
		"code": 401,
		"data": data,
		"msg":  msg,
	})
}
