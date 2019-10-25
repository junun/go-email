package main

import (
	"fmt"
	"go_email/src/pkg/setting"
	"go_email/src/go_email"
	"go_email/src/routers"
	"net/http"
)


func main() {
	// 主进程运行期间启动一个定时任务协程
	go func() {
		go_email.SendEmailCronTask()
	}()


	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        r,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
