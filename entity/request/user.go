// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/9 15:03

package request

import uuid "github.com/satori/go.uuid"

type Login struct {
	Username  string `json:"username" binding:"required" ` // 用户名
	Password  string `json:"password" binding:"required"`  // 密码
	Captcha   string `json:"captcha" binding:"required"`   // 验证码
	CaptchaId string `json:"captchaId" binding:"required"` // 验证码ID
}

type UserList struct {
	Id       string   `json:"id"`       // id
	Ids      []string `json:"ids"`      // ids
	NickName string   `json:"nickName"` // 昵称
	Phone    string   `json:"phone"`    // 电话号码
	Email    string   `json:"email"`    // 邮箱
	Pages
}

type UpdateEnable struct {
	Id     string `json:"id" binding:"required"`     // id
	Enable int    `json:"enable" binding:"required"` // 开关
}

type Update struct {
	Id       string `json:"id" binding:"required" db:"-"` // id
	NickName string `json:"nickName" db:"nick_name"`      // 昵称
	Phone    string `json:"phone" db:"phone"`             // 电话
	Email    string `json:"email" db:"email"`             // 邮箱
}

type Create struct {
	NickName    string    `json:"nickName" db:"nick_name"`       // 昵称
	Phone       string    `json:"phone" db:"phone"`              // 电话
	Email       string    `json:"email" db:"email"`              // 邮箱
	UUID        uuid.UUID `json:"uuid" db:"uuid"`                // uuid
	UserName    string    `json:"userName"db:"username"`         // 用户名
	Password    string    `json:"password" db:"password"`        // 密码
	SideMode    string    `json:"sideMode" db:"side_mode"`       // 用户侧边主题
	HeaderImg   string    `json:"headerImg" db:"header_img"`     // 用户头像
	BaseColor   string    `json:"baseColor" db:"base_color"`     // 基础颜色
	ActiveColor string    `json:"activeColor" db:"active_color"` // 活跃颜色
	AuthorityId int       `json:"authorityId" db:"authority_id"` // 用户角色ID
}

type Delete struct {
	Id string `json:"id" binding:"required"`
}
