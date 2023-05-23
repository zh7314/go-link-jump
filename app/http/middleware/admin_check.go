package middleware

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/services"
	"goravel/app/utils"
	"goravel/app/utils/global"
	"goravel/app/utils/response"
)

func AdminCheck() http.Middleware {
	return func(ctx http.Context) {

		_, err := check(ctx)
		if err != nil {
			response.AbortFail(ctx, "", err.Error())
		}

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("系统内部错误")
				response.Fail(ctx, "", utils.ErrorToString(r))
				return
			}
		}()

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
	//fmt.Println(token)

	apiService := services.NewApiService()

	key := apiService.GetTokenKey() + token

	has := facades.Cache.Has(key)
	if !has {
		return false, errors.New("查询api key不存在")
	}

	value := facades.Cache.GetString(key)
	//fmt.Println(value)

	if len(value) == 0 {
		return false, errors.New("api key不存在")
	}

	ctx.WithValue(apiService.GetApiKey(), value)

	return true, nil
}
