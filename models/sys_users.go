// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/9 14:46

package models

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/constant"
	"github.com/EDDYCJY/go-gin-example/entity"
	"github.com/EDDYCJY/go-gin-example/entity/request"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/pkg/util/xstruct"
	"strings"
)

type SystemUser struct {
}

// GetTableName
//
//  @Description:  获取数据表名称
//  @return string
//
func (su *SystemUser) GetTableName() string {
	return entity.SysUser{}.TableName()
}

// GetInfoByName
//
//  @Description:  查name获取信息
//  @param u
//  @return *entity.SysUser
//  @return error
//
func (su *SystemUser) GetInfoByName(u *entity.SysUser) (*entity.SysUser, error) {
	var userInfo entity.SysUser
	err := db.Where("username = ?", u.Username).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

// List
//
//  @Description: 获取用户列表
//  @param req
//  @return []*entity.SysUser
//  @return error
//
func (su *SystemUser) List(req request.UserList) ([]*entity.SysUser, error) {
	var (
		sql  strings.Builder
		args []interface{}
	)
	rsp := make([]*entity.SysUser, 0)
	table := su.GetTableName()
	sql.WriteString("SELECT * FROM " + table + " WHERE 1=1 ")
	if req.Id != "" {
		sql.WriteString(" AND `id` = ?")
		args = append(args, req.Id)
	}
	if req.Email != "" {
		sql.WriteString(" AND `email` = ?")
		args = append(args, req.Email)
	}
	if req.NickName != "" {
		sql.WriteString(" AND `nick_name` like ?")
		args = append(args, "%"+req.NickName+"%")
	}

	if len(req.Ids) > 0 {
		sql.WriteString(" AND `id` IN (?)")
		args = append(args, req.Ids)
	}

	sql.WriteString(" ORDER BY id DESC ")
	offset, limit := util.GetPageLimitInt(req.Page, req.Limit)
	sql.WriteString(" LIMIT %d, %d")
	err := db.Table(table).Raw(fmt.Sprintf(sql.String(), offset, limit), args...).Find(&rsp).Error
	return rsp, err
}

// SetEnable
//
//  @Description: 更新enable开关
//  @param id
//  @param enable
//  @return bool
//
func (su *SystemUser) SetEnable(id string, enable int) bool {
	table := su.GetTableName()
	err := db.Table(table).Where("id = ?", id).Update("enable", enable).Error
	if err != nil {
		return false
	}

	return true
}

// Update
//
//  @Description: 更新信息
//  @param req
//  @return int64
//  @return error
//
func (su *SystemUser) Update(req request.Update) (int64, error) {
	var (
		updateSql strings.Builder
		args      []interface{}
		data      = map[string]interface{}{}
	)
	err := xstruct.Decode(req, &data, xstruct.WithTagName(constant.DBTag))
	if err != nil {
		return 0, err
	}
	table := su.GetTableName()
	updateSql.WriteString("UPDATE " + table + " SET ")
	placeArgs := make([]string, 0, len(data))
	for k, v := range data {
		placeArgs = append(placeArgs, k+"=?")
		args = append(args, v)
	}
	updateSql.WriteString(strings.Join(placeArgs, ","))
	updateSql.WriteString(" WHERE `id` = ?")
	args = append(args, req.Id)

	effects := db.Exec(updateSql.String(), args...)
	return effects.RowsAffected, effects.Error
}

// Create
//
//  @Description: 创建用户
//  @param req
//  @return error
//
func (su *SystemUser) Create(req request.Create) error {
	ret := entity.SysUser{}
	err := xstruct.Decode(req, &ret, xstruct.WithTagName(constant.JsonTag))
	if err != nil {
		return err
	}
	ret.Password = req.Password
	sql := fmt.Sprintf("INSERT INTO %s (`username`,`password`,`nick_name`,`side_mode`,`header_img`,`base_color`,`active_color`,`authority_id`,`phone`,`email`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?)", su.GetTableName())
	err = db.Table(su.GetTableName()).Raw(sql, ret.Username, ret.Password, ret.NickName, ret.SideMode, ret.HeaderImg, ret.BaseColor, ret.ActiveColor, ret.AuthorityId, ret.Phone, ret.Email).Create(&ret).Error
	return err
}

// Delete
//
//  @Description: 删除数据
//  @receiver su
//  @param id
//  @return error
//
func (su *SystemUser) Delete(id string) error {
	sql := "DELETE FROM " + su.GetTableName() + " WHERE `id` = ?"
	effects := db.Exec(sql, id)
	return effects.Error
}

// Count
//
//  @Description: 查询条目数
//  @receiver su
//  @param req
//  @return int64
//  @return error
//
func (su *SystemUser) Count(req request.UserList) (int64, error) {
	var (
		sql  strings.Builder
		args []interface{}
	)
	var total int64
	table := su.GetTableName()
	sql.WriteString("SELECT count(*) FROM " + table + " WHERE 1=1 ")
	if req.Id != "" {
		sql.WriteString(" AND `id` = ?")
		args = append(args, req.Id)
	}
	if req.Email != "" {
		sql.WriteString(" AND `email` = ?")
		args = append(args, req.Email)
	}
	if req.NickName != "" {
		sql.WriteString(" AND `nick_name` like ?")
		args = append(args, "%"+req.NickName+"%")
	}

	if len(req.Ids) > 0 {
		sql.WriteString(" AND `id` IN (?)")
		args = append(args, req.Ids)
	}

	err := db.Table(table).Raw(sql.String(), args...).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
