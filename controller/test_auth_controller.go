// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/10 09:57

package controller

import (
	"github.com/EDDYCJY/go-gin-example/utils"
	"github.com/gin-gonic/gin"
)

func TestAuth(c *gin.Context) {

	utils.OkWithDetailed("dsd", "", c)
}
