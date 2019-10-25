package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File
	RunMode string
	AppSetting = &App{}
	ServerSetting = &Server{}
)

type App struct {
	JwtSecret string
	PageSize int
	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	if RunMode == "debug" {
		Cfg, err = ini.Load("conf/debug.ini")
		if err != nil {
			log.Fatalf("Fail to parse 'conf/debug.ini': %v", err)
		}
	} else {
		Cfg, err = ini.Load("conf/release.ini")
		if err != nil {
			log.Fatalf("Fail to parse 'conf/release.ini': %v", err)
		}
	}

	LoadServer()
	LoadApp()
}


func LoadServer() {
	err := Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

func LoadApp() {
	err := Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
}

