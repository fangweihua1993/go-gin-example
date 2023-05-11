// @Description
// @Author yixia
// @Copyright 2021 sndks.com. All rights reserved.
// @LastModify 2021/1/14 5:21 下午

package errs

import "github.com/EDDYCJY/go-gin-example/utils"

var UserNotFound = utils.NewCustomError(10001, "用户未找到")
var PasswordError = utils.NewCustomError(10002, "密码错误")

var CaptchaError = utils.NewCustomError(11001, "验证码创建失败")
var CaptchaCheckError = utils.NewCustomError(11002, "验证码验证失败")

var SystemError = utils.NewCustomError(500, "系统错误")
var ParamError = utils.NewCustomError(1, "参数错误")

var MysqlSelectError = utils.NewCustomError(100, "查询失败")
var MysqlUpdateError = utils.NewCustomError(101, "更新失败")
var MysqlCreateError = utils.NewCustomError(102, "创建失败")
var MysqlDeleteError = utils.NewCustomError(103, "删除失败")
