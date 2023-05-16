package middleware

import (
	"github.com/goravel/framework/contracts/http"
)

func AdminCheck() http.Middleware {
	return func(ctx http.Context) {

		//ctx.WithValue("user", "Goravel")

		ctx.Request().Next()
	}
}
