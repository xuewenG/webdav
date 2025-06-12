package router

import (
	"net/http"

	"github.com/xuewenG/webdav/pkg/handler"
	"github.com/xuewenG/webdav/pkg/util"
)

func NewRouter() http.Handler {
	// 创建路由实例
	mux := http.NewServeMux()

	// 添加健康检查路由
	mux.HandleFunc(util.WithPrefix("/health"), handler.HealthCheckHandler)
	// 添加 WebDAV 路由
	mux.Handle(util.WithPrefix("/dav/"), handler.NewWebDAVHandler())

	return mux
}
