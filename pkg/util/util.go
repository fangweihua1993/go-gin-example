package util

import "github.com/EDDYCJY/go-gin-example/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

func GetPageLimitInt(pageNo int, pageSize int) (int, int) {
	pageNo, pageSize = GetPageAndSize(pageNo, pageSize)
	return (pageNo - 1) * pageSize, pageSize
}

func GetPageAndSize(pageNo int, pageSize int) (int, int) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 || pageSize > 10000 {
		pageSize = 20
	}
	return pageNo, pageSize
}
