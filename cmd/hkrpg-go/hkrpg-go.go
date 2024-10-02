package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gucooing/hkrpg-go/hkrpg-go-pe"
	"github.com/gucooing/hkrpg-go/pkg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/push/client"
)

func main() {
	confName := "hkrpg-go-pe.json"
	err := hkrpg_go_pe.LoadConfig(confName)
	if err != nil {
		if err == hkrpg_go_pe.FileNotExist {
			p, _ := json.MarshalIndent(hkrpg_go_pe.DefaultConfig, "", "  ")
			cf, _ := os.Create("./conf/" + confName)
			cf.Write(p)
			cf.Close()
			fmt.Printf("找不到配置文件\n已生成默认配置文件 %s \n", confName)
			main()
		} else {
			panic(err)
		}
	}
	cfg := hkrpg_go_pe.GetConfig()
	// 初始化日志
	logger.InitLogger("hkrpg_go_pe", strings.ToUpper(cfg.LogLevel))
	logger.Info("hkrpg_go_pe")
	logger.Info("AppVersion:%s", pkg.GetAppVersion())
	logger.Info("GameVersion:%s", pkg.GetGameVersion())
	client.NewPushClient(cfg.PushUrl)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// 初始化服务器
	s := hkrpg_go_pe.NewServer(cfg)

	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			logger.Info("hkrpg_go_pe服务正在关闭")
			s.Close()
			logger.Info("hkrpg_go_pe服务已停止")
			logger.CloseLogger()
			os.Exit(0)
		}
	}()

	select {}
}
