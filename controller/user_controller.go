// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/9 16:20

package controller

import (
	"github.com/EDDYCJY/go-gin-example/entity/request"
	"github.com/EDDYCJY/go-gin-example/service"
	"github.com/EDDYCJY/go-gin-example/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req request.Login
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret, err := service.Login(c, req)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}

func GetUserList(c *gin.Context) {
	var req request.UserList
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret, err := service.GetUserList(req)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}

func UpdateEnable(c *gin.Context) {
	var req request.UpdateEnable
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret := service.UpdateEnable(req)

	utils.OkWithDetailed(ret, "", c)
}

func Update(c *gin.Context) {
	var req request.Update
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret, err := service.Update(req)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}

func Create(c *gin.Context) {
	var req request.Create
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret, err := service.Create(req)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}

func Delete(c *gin.Context) {
	var req request.Delete
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret, err := service.Delete(req)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}
