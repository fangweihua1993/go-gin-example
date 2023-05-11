// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/11 14:15

package request

// Pages 公共入参
type Pages struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}
