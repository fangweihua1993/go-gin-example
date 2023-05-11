// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/9 15:03

package service

import (
	"fmt"
	errs "github.com/EDDYCJY/go-gin-example/eerrs"
	"github.com/EDDYCJY/go-gin-example/entity"
	"github.com/EDDYCJY/go-gin-example/entity/request"
	"github.com/EDDYCJY/go-gin-example/entity/response"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func Login(c *gin.Context, req request.Login) (*response.Login, error) {
	var ret response.Login
	// 判断验证码是否开启
	key := c.ClientIP()
	openCaptcha := setting.AppSetting.CaptchaOpen           // 是否开启防爆次数
	openCaptchaTimeOut := setting.AppSetting.CaptchaTimeout // 缓存超时时间
	v, err := gredis.Get(key)
	if err != nil && err.Error() != "redigo: nil returned" {
		logging.Error(fmt.Printf("[captcha_service.Captcha.redis]redis获取失败，err：%+v", err))
		return nil, errs.SystemError
	}
	if v == nil {
		gredis.Set(key, 1, openCaptchaTimeOut)
	}

	var oc bool = openCaptcha == 0 || openCaptcha < utils.ToInt(v)

	if !oc || store.Verify(req.CaptchaId, req.Captcha, true) {
		userLogin := entity.SysUser{
			Username: req.Username,
		}
		userInfo, err := models.GetInfoByName(&userLogin)
		if err != nil || userInfo == nil {
			return nil, errs.UserNotFound
		}
		if userInfo.Enable != 1 {
			logging.Error("登陆失败! 用户名不存在或者密码错误!")
			return nil, err
		}
		if userInfo != nil && userInfo.Username != "" {
			if ok := utils.BcryptCheck(req.Password, userInfo.Password); !ok {
				return nil, errs.PasswordError
			}
			ret.Jwt, _ = util.GenerateToken(userInfo.Username, userInfo.Password)
			ret.UserId = utils.ToString(userInfo.ID)

			return &ret, nil
		}
	}

	return &ret, errs.CaptchaCheckError
}

func GetUserList(req request.UserList) ([]*entity.SysUser, error) {
	return models.List(req)
}

func UpdateEnable(req request.UpdateEnable) bool {
	return models.SetEnable(req.Id, req.Enable)
}

func Update(req request.Update) (bool, error) {
	_, err := models.Update(req)
	if err != nil {
		return false, errs.MysqlUpdateError
	}

	return true, nil
}

func Create(req request.Create) (bool, error) {
	req.Password = utils.BcryptHash(req.Password)
	req.UUID = uuid.NewV4()
	err := models.Create(req)
	if err != nil {
		return false, errs.MysqlCreateError
	}

	return true, nil
}

func Delete(req request.Delete) (bool, error) {
	err := models.Delete(req.Id)
	if err != nil {
		return false, errs.MysqlDeleteError
	}

	return true, nil
}
