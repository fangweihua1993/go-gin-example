// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/9 14:46

package models

import "github.com/EDDYCJY/go-gin-example/entity"

func Login(u *entity.SysUser) (*entity.SysUser, error) {
	var userInfo entity.SysUser
	err := db.Where("username = ?", u.Username).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}
