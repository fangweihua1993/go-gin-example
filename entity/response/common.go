// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/16 14:10

package response

//ListData 分页数据
type ListData struct {
	//Page  当前数据所在页
	Page int `json:"page"`
	//Limit 每页显示条数
	Limit int `json:"limit"`
	//Count 本次返回数据条数
	Count int `json:"count"`
	//Cursor 游标值
	Cursor int64 `json:"cursor"`
	//Total 服务器估算数据总条数
	Total int64 `json:"total"`
	//List 数组列表
	List interface{} `json:"list"`
}
