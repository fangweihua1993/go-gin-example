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

type UserController struct {
	userService service.UserService
}

func (u *UserController) Login(c *gin.Context) {
	var req request.Login
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret, err := u.userService.Login(c, req)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}

func (u *UserController) GetUserList(c *gin.Context) {
	var req request.UserList
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret, err := u.userService.GetUserList(req)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}

func (u *UserController) UpdateEnable(c *gin.Context) {
	var req request.UpdateEnable
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret := u.userService.UpdateEnable(req)

	utils.OkWithDetailed(ret, "", c)
}

func (u *UserController) Update(c *gin.Context) {
	var req request.Update
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret, err := u.userService.Update(req)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}

func (u *UserController) Create(c *gin.Context) {
	var req request.Create
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret, err := u.userService.Create(req)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}

func (u *UserController) Delete(c *gin.Context) {
	var req request.Delete
	if err := c.ShouldBind(&req); err != nil {
		utils.FailWithDetailed(nil, "参数错误", c)
		return
	}

	ret, err := u.userService.Delete(req)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}
