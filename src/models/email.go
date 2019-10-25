package models

import (
	"go_email/src/pkg/setting"
	"gopkg.in/gomail.v2"
	"log"
)

var (
	Gmd  *gomail.Dialer
)

func init() {
	// 获取 email 配置
	e, err := setting.Cfg.GetSection("email")
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis': %v", err)
	}

	port, _ := e.Key("port").Int()

	Gmd = InitDialer(
		e.Key("host").String(),
		e.Key("user").String(),
		e.Key("pass").String(),
		port)
}

// 自定义发送邮箱
func InitDialer(host, user, pass string , port int) *gomail.Dialer {
	return gomail.NewDialer(host, port, user, pass)
}
