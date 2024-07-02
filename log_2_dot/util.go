package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func marshal2String(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}
	return string(bytes)
}

func unmarshal2map(in string) map[string]interface{} {
	out := make(map[string]interface{})
	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		return nil
	}
	return out
}

func getIpv4() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("Main IPv4 Address:", ipNet.IP.String())
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
