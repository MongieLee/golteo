package service

import (
	"ginl/utils"
	"github.com/go-co-op/gocron/v2"
	"time"
)

func StartJob() {
	// 创建一个调度器
	s, err := gocron.NewScheduler()
	if err != nil {

	}
	// 给调度器添加任务
	_, err = s.NewJob(
		gocron.DurationJob(10*time.Second),
		gocron.NewTask(func(a string, b int) {
			utils.InfoF("test job run")
		},
			"test", 1))
	if err != nil {

	}
	s.Start()
	select {
	case <-time.After(time.Second * 20):
		err := s.Shutdown()
		if err != nil {

		}
		utils.InfoF("停止定时任务")
	}
}
