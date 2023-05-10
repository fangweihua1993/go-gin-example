// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/9 18:59

package controller

import (
	"github.com/EDDYCJY/go-gin-example/service"
	"github.com/EDDYCJY/go-gin-example/utils"
	"github.com/gin-gonic/gin"
)

func Captcha(c *gin.Context) {

	ret, err := service.Captcha(c)
	if err != nil {
		utils.FailWithDetailed(nil, err.Error(), c)
		return
	}

	utils.OkWithDetailed(ret, "", c)
}
