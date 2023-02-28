package handle

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/luoruofeng/DockerApiAgent/model"
	"github.com/luoruofeng/DockerApiAgent/util"
	"go.uber.org/zap"
)

func AgentFunc(c *http.Client, logger *zap.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 将请求转发到目标 HTTP 客户端
		r.URL.Path = r.URL.Path[len(model.Cnf.AgentPathPrefix):]
		r.RequestURI = ""
		r.URL.Scheme = "http"
		r.URL.Host = "localhost"
		util.LogInfo(logger, fmt.Sprintf("request docker api. request parameter:%v", r))
		resp, err := c.Do(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// 将响应正文返回给客户端
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(body)
	})
}
