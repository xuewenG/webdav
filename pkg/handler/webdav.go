package handler

import (
	"log"
	"net/http"

	"github.com/xuewenG/webdav/pkg/middleware"
	"golang.org/x/net/webdav"
)

// NewWebDAVHandler 创建 WebDAV 处理器
func NewWebDAVHandler(prefix string, rootDir string, readOnly bool) http.Handler {
	// 创建处理器
	var handler http.Handler
	handler = &webdav.Handler{
		Prefix:     prefix,
		FileSystem: webdav.Dir(rootDir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err == nil {
				log.Printf("WebDAV, username: %s, method: %s, path: %s", middleware.GetUsername(r), r.Method, r.URL.Path)
			} else {
				log.Printf("WebDAV, username: %s, method: %s, path: %s, error: %v", middleware.GetUsername(r), r.Method, r.URL.Path, err)
			}
		},
	}

	// 使用只读中间件
	if readOnly {
		handler = middleware.ReadOnlyMiddleware(handler)
	}

	// 使用认证中间件
	handler = middleware.AuthMiddleware(handler)

	return handler
}
