// @Description
// @Author fangweihua@yixia.com
// @Copyright 2023 sndks.com. All rights reserved.
// @Datetime 2023/5/16 14:49

package task

import (
	"github.com/EDDYCJY/go-gin-example/entity/request"
	"github.com/EDDYCJY/go-gin-example/service"
	"github.com/robfig/cron"
	"time"
)

type UserListTask struct {
	userService service.UserService
}

// UserListTimer
//
//  @Description: 定时任务
//  @receiver u
//
func (u *UserListTask) UserListTimer() {
	c := cron.New()
	// 每5s执行一次 执行中的停止定时 执行完继续定时 避免未执行完重复执行
	c.AddFunc("*/5 * * * *", func() {
		c.Stop()
		time.Sleep(10 * time.Duration(time.Second))
		var req request.UserList
		u.userService.GetUserList(req)
		c.Start()
	})
	c.Start()
}
