package utils

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"hash/crc32"
	"strings"
)

func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func CRC32(input string) uint32 {
	bytes := []byte(input)
	return crc32.ChecksumIEEE(bytes)
}

func GetIP(ctx http.Context) string {

	ip := strings.TrimSpace(ctx.Request().Header("X-Forwarded-For", "127.0.0.1"))
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(ctx.Request().Header("X-Real-Ip", "127.0.0.1"))
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(ctx.Request().Ip())
	if ip != "" {
		return ip
	}
	return "127.0.0.1"
}

// 检索顺序为：json, form, query, route
func GetAllRequestParameters(ctx http.Context) {

	//ctx.Request().QueryMap()
	//ctx.Request().Input()

	fmt.Println(ctx.Value("zx"))
}
