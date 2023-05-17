package middleware

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
)

func ApiLog() http.Middleware {
	return func(ctx http.Context) {

		method := ctx.Request().Method()
		//url := ctx.Request().Url()
		//path := ctx.Request().Path()
		//ip := utils.GetIP(&ctx)

		// 目前系统基本只支持Post Get Options
		if method == http.MethodGet {

		} else if method == http.MethodPost {

		} else {
			fmt.Println(method)
		}

		//fmt.Println(method)
		//fmt.Println(url)
		//fmt.Println(path)
		//fmt.Println(ip)

		fmt.Println(ctx.Request().Origin())

		fmt.Println(ctx.Request().Origin())
		fmt.Println("-----------------------------------------\n")
		fmt.Println(ctx.Response().Origin().Header())
		fmt.Println(ctx.Response().Origin().Body())

		ctx.Request().Next()

	}
}
