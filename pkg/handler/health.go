package handler

import (
	"net/http"
)

// HealthCheckHandler 处理健康检查请求
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
