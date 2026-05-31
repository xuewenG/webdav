package middleware

import (
	"log"
	"net/http"
)

// ReadOnlyMiddleware 限制只允许读取操作
func ReadOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 只允许 GET、HEAD、OPTIONS、PROPFIND 方法
		switch r.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions, "PROPFIND":
			next.ServeHTTP(w, r)
			return
		default:
			log.Printf("ReadOnlyMiddleware, username: %s, method: %s, path: %s", GetUsername(r), r.Method, r.URL.Path)
			http.Error(w, "只读模式, 不允许写入操作", http.StatusForbidden)
		}
	})
}
