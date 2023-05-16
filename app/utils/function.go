package utils

import (
	"github.com/goravel/framework/contracts/http"
	"hash/crc32"
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

func GetIP(ctx *http.Context) string {

	return ""
}