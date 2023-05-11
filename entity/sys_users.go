// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/9 14:40

package entity

import uuid "github.com/satori/go.uuid"

type SysUser struct {
	ID          int       `json:"id" gorm:"primarykey"`                                                                 // 主键ID
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                                                     // 用户UUID
	Username    string    `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	SideMode    string    `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg   string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string    `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor string    `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId uint      `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Phone       string    `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	Email       string    `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	Enable      int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                      //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}
