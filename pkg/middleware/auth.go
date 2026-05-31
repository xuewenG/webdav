package middleware

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/xuewenG/webdav/pkg/config"
)

// 定义 context key
type contextKey string

const UsernameKey contextKey = "username"

// GetUsername 从 context 中读取用户名
func GetUsername(r *http.Request) string {
	if username, ok := r.Context().Value(UsernameKey).(string); ok {
		return username
	}

	return ""
}

// AuthMiddleware 验证用户登录信息
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取 Authorization 头
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="WebDAV"`)
			http.Error(w, "需要认证", http.StatusUnauthorized)
			return
		}

		// 解析 Basic 认证信息
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			http.Error(w, "无效的认证格式", http.StatusBadRequest)
			return
		}

		// 解码认证信息
		payload, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			http.Error(w, "无效的认证信息", http.StatusBadRequest)
			return
		}

		// 解析用户名和密码
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "无效的认证信息", http.StatusBadRequest)
			return
		}

		username, password := pair[0], pair[1]

		// 验证用户
		authenticated := false
		for _, user := range config.Config.Users {
			if user.Username == username && user.Password == password {
				authenticated = true
				break
			}
		}

		if !authenticated {
			w.Header().Set("WWW-Authenticate", `Basic realm="WebDAV"`)
			http.Error(w, "认证失败", http.StatusUnauthorized)
			return
		}

		// 将用户名保存到 context 中
		ctx := context.WithValue(r.Context(), UsernameKey, username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
