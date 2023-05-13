package middleware

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
)

func Recovery() http.Middleware {
	return func(ctx http.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("系统内部错误")

				ctx.Response().Json(http.StatusOK, http.Json{
					"msg": r,
				})
			}
		}()

		ctx.Request().Next()
	}
}
