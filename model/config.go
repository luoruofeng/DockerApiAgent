package model

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var Cnf *Config

func CreateConfig(configFile string) {
	//default value
	if configFile == "" {
		absPath, err := filepath.Abs(os.Args[0])
		if err != nil {
			panic(err)
		}
		// 获取执行程序所在的目录
		dir := filepath.Dir(absPath)
		// 拼接文件路径
		configFile = filepath.Join(dir, "./config.json")
	}

	// 加载配置
	cbs, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(cbs, &Cnf)
	if err != nil {
		log.Fatalf("failed to load configuration: %s", err)
	}
}

type Config struct {
	ConsulAddr           string `json:"consul_addr"`
	ConsulHealth         bool   `json:"consul_health"`
	ConsulHealthInterval string `json:"consul_health_interval"`
	ConsulHealthTimeout  string `json:"consul_health_timeout"`
	ConsulHealthPort     string `json:"consul_health_port"`
	NICName              string `json:"nic_name"`
	LogLevel             string `json:"log_level"`
	LogFile              string `json:"log_file"`
	HttpAddr             string `json:"http_addr"`
	UnixFile             string `json:"unix_file"`
	ServiceName          string `json:"service_name"`
	HttpReadOverTime     int    `json:"http_read_over_time"`
	HttpWriteOverTime    int    `json:"http_write_over_time"`
	AgentPathPrefix      string `json:"agent_path_prefix"`
	IsProduction         bool   `json:"is_production"`
}
