package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/HarryLuo227/simple-blog-service/global"
	"github.com/HarryLuo227/simple-blog-service/internal/model"
	"github.com/HarryLuo227/simple-blog-service/internal/routers"
	"github.com/HarryLuo227/simple-blog-service/pkg/logger"
	"github.com/HarryLuo227/simple-blog-service/pkg/setting"
	"github.com/HarryLuo227/simple-blog-service/pkg/tracer"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	port         string
	runMode      string
	config       string
	isVerson     bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {
	setupFlag()

	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Jaeger", &global.JaegerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}

	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", global.JaegerSetting.AgentHost)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer

	return nil
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "啟動通訊埠")
	flag.StringVar(&runMode, "mode", "", "啟動模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的設定檔路徑")
	flag.BoolVar(&isVerson, "version", false, "編譯資訊")
	flag.Parse()

	return nil
}

// @Title 部落格系統
// @Version 1.0
// @Description Go - Simple-Blog-Service
func main() {
	if isVerson {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
		return
	}

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %s", err)
		}
	}()

	// 等待中斷訊號
	quit := make(chan os.Signal)
	// 接受 syscall.SIGINT 和 syscall.SIGTERM 訊號
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 最大時間控制，用於通知該服務端它有 5 秒的時間來處理原有的請求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shut down:", err)
	}

	log.Println("Server exiting")
}
