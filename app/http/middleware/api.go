package middleware

import (
	"github.com/goravel/framework/contracts/http"
)

func Api() http.Middleware {
	return func(ctx http.Context) {
		ctx.Request().Next()
	}
}
