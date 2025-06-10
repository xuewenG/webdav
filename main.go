package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xuewenG/webdav/pkg/config"
	"github.com/xuewenG/webdav/pkg/handler"
)

func main() {
	// 初始化配置
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
		return
	}

	// 获取配置
	port := config.Config.Port
	prefix := config.Config.Prefix
	rootDir := config.Config.RootDir
	readOnly := config.Config.ReadOnly
	log.Printf("port: %d\nprefix: %s\nrootDir: %s\nreadOnly: %t", port, prefix, rootDir, readOnly)

	// 确保根目录存在
	if err := os.MkdirAll(rootDir, 0755); err != nil {
		log.Fatal(err)
		return
	}

	// 初始化路由
	mux := http.NewServeMux()
	// 添加健康检查路由
	mux.HandleFunc("/health", handler.HealthCheckHandler)
	// 添加 WebDAV 路由
	mux.Handle(fmt.Sprintf("%s/", prefix), handler.NewWebDAVHandler(prefix, rootDir, readOnly))

	// 启动服务器
	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
