package middleware

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils"
)

func Recovery() http.Middleware {
	return func(ctx http.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("系统内部错误")
				utils.Fail(ctx, "", utils.ErrorToString(r))
			}
		}()

		ctx.Request().Next()
	}
}
