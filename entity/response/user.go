// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/9 15:07

package response

type Login struct {
	Jwt    string `json:"jwt"`
	UserId string `json:"userId"`
}
