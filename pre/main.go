package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	// 创建 HTTP 客户端
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial("unix", "/var/run/docker.sock")
		},
	}
	client := &http.Client{Transport: transport}

	// 创建代理服务器
	proxy := &http.Server{
		Addr: ":8888",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("%v\n", r)
			// 将请求转发到目标 HTTP 客户端
			r.RequestURI = ""
			r.URL.Scheme = "http"
			r.URL.Host = "localhost"
			resp, err := client.Do(r)
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
		}),
	}

	// 启动代理服务器
	fmt.Println("Proxy listening on", proxy.Addr)
	if err := proxy.ListenAndServe(); err != nil {
		panic(err)
	}
}
