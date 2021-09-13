package models

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/robfig/cron/v3"
	"math/rand"
)

func intiSky() {
	c := cron.New(cron.WithSeconds()) //精确到秒

	//定时任务
	spec := "0 " + string(rand.Intn(59)) + " " + Config.CTime + "/12 * * ?" //cron表达式，每秒一次
	logs.Info(spec)
	c.AddFunc(spec, func() {
		fmt.Println("开始wskey转换")
		updateCookie()
	})

	c.Start()
}