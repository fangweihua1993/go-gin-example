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
