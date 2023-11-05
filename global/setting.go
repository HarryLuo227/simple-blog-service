package global

import (
	"github.com/HarryLuo227/simple-blog-service/pkg/logger"
	"github.com/HarryLuo227/simple-blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	JaegerSetting   *setting.JaegerSetting
	JWTSetting      *setting.JWTSetting
	Logger          *logger.Logger
)
