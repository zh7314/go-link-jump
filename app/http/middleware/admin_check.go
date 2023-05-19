package middleware

import (
	"errors"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils/global"
)

func AdminCheck() http.Middleware {
	return func(ctx http.Context) {

		//ctx.WithValue("user", "Goravel")

		//method := ctx.Request().Method()

		//var token string

		//if method == http.MethodPost {
		//
		//
		//} else if method == http.MethodGet {
		//
		//} else {
		//	fmt.Println(method)
		//}

		ctx.Request().Next()
	}
}
func check(ctx http.Context) (res bool, ok error) {
	token := ctx.Request().Header(global.API_TOKEN, "")
	if token == "" {
		token = ctx.Request().Input(global.API_TOKEN)
	}

	if token == "" {
		return false, errors.New("token不能为空")
	}

	return true, nil
}
