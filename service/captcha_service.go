// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/9 17:40

package service

import (
	"fmt"
	errs "github.com/EDDYCJY/go-gin-example/eerrs"
	"github.com/EDDYCJY/go-gin-example/entity/response"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) (*response.SysCaptchaResponse, error) {
	// 判断验证码是否开启
	openCaptcha := setting.AppSetting.CaptchaOpen           // 是否开启防爆次数
	openCaptchaTimeOut := setting.AppSetting.CaptchaTimeout // 缓存超时时间
	key := c.ClientIP()
	v, err := gredis.Get(key)
	if err != nil && err.Error() != "redigo: nil returned" {
		logging.Error(fmt.Printf("[captcha_service.Captcha.redis]redis获取失败，err：%+v", err))
		return nil, errs.SystemError
	}
	if v == nil {
		gredis.Set(key, 1, openCaptchaTimeOut)
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < utils.ToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(setting.AppSetting.CaptchaImgHeight, setting.AppSetting.CaptchaImgWidth, setting.AppSetting.CaptchaNumber, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		logging.Error(fmt.Printf("[captcha_service.Captcha.Generate]验证码获取失败，err：%+v", err))
		return nil, errs.CaptchaError
	}
	ret := response.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: setting.AppSetting.CaptchaNumber,
		OpenCaptcha:   oc,
	}
	return &ret, nil
}
