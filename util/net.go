package util

import (
	"net"
)

// 根据网卡接口名获取对应的ipv4的地址
func GetIpByNICName(name string) string {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		panic(err)
	}
	addrs, err := iface.Addrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return ""
}
