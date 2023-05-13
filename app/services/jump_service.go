package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	"hash/crc32"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const jumpKey = "JumpLink:"

type JumpService struct {
	//Dependent services
	Key string
}

func NewJumpService() *JumpService {
	return &JumpService{
		//Inject model
		Key: jumpKey,
	}
}

func (r *JumpService) GetData() []models.Jump {
	var jump []models.Jump
	facades.Orm.Query().Get(&jump)

	return jump
}

func (r *JumpService) DoJump(id string) (jumpUrl string, ok error) {

	if len(strings.TrimSpace(id)) == 0 {
		return "", errors.New("查询条件不能为空")
	}

	key := r.Key + id
	isHas := facades.Cache.Has(key)
	if isHas {
		return "", errors.New("查询key存在")
	}

	res := facades.Cache.GetString(key)
	if len(res) == 0 {
		return "", errors.New("查询值为空")
	}

	jump, err := url.QueryUnescape(res)
	if err != nil {
		return "", err
	}

	return jump, nil
}

func (r *JumpService) AddLink(url_str string, end_time string) (res bool, ok error) {

	if len(strings.TrimSpace(url_str)) == 0 {
		return false, errors.New("url不能为空")
	}
	if len(strings.TrimSpace(end_time)) == 0 {
		return false, errors.New("end_time不能为空")
	}

	now := time.Now()
	start_time := now.Unix()

	end, err := time.Parse("2006-01-02 15:04:05", end_time)
	if err != nil {
		return false, errors.New("解析结束时间错误")
	}

	end_time_unix := end.Unix()
	if start_time > end_time_unix {
		return false, errors.New("结束时间不能大于当前时间")
	}

	count := end_time_unix - start_time

	var jump models.Jump
	jump.JumpUrl = url.QueryEscape(url_str)
	jump.EndTime = end

	err1 := facades.Orm.Query().Save(&jump)
	if err1 != nil {
		return false, errors.New("保存失败")
	}

	str_id := strconv.FormatInt(int64(jump.ID), 10)
	m_id := CRC32(str_id)
	jump.Md5Data = strconv.FormatInt(int64(m_id), 10)

	err2 := facades.Orm.Query().Save(&jump)
	if err2 != nil {
		return false, err2
	}

	key := r.Key + jump.Md5Data
	err3 := facades.Cache.Put(key, jump.JumpUrl, time.Duration(count)*time.Second)
	if err3 != nil {
		return false, err3
	}
	
	return true, ok
}

func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

func GeneratorMD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}

func CRC32(input string) uint32 {
	bytes := []byte(input)
	return crc32.ChecksumIEEE(bytes)
}
