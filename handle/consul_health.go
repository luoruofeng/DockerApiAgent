package handle

import (
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	// 检查服务的状态，返回对应的HTTP响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
