package bootstrap

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
	"xiaozhuquan.com/xiaozhuquan/global"
)

func InitializeCron() {
	global.App.Cron = cron.New(cron.WithSeconds())

	go func() {
		_, _ = global.App.Cron.AddFunc("0 0 2 * * *", func() {
			fmt.Println(time.Now())
		})
		global.App.Cron.Start()
		defer global.App.Cron.Stop()
		select {}
	}()
}
