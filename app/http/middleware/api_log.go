package middleware

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils"
)

func ApiLog() http.Middleware {
	return func(ctx http.Context) {

		method := ctx.Request().Method()
		url := ctx.Request().Url()
		path := ctx.Request().Path()
		ip := utils.GetIP(&ctx)

		ctx.Request().Origin()

		// 目前系统基本只支持Post Get Options
		if method == http.MethodGet {

		} else if method == http.MethodPost {

		} else {
			fmt.Println(method)
		}

		fmt.Println(method)
		fmt.Println(url)
		fmt.Println(path)
		fmt.Println(ip)

		ctx.Request().Next()

		ctx.Response().Success()
	}
}
