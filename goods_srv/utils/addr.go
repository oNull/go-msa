package utils

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

const (
	minPort     = 40000
	maxPort     = 50000
	maxAttempts = 100 // 最大尝试次数
)

func GetFreePort() (int, error) {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	_, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("localhost:%d", minPort))
	if err != nil {
		return 0, err
	}

	// 获取在指定范围内的随机空闲端口
	port, err := getFreePortInRange(minPort, maxPort)
	if err != nil {
		return 0, err
	}

	// 启动服务并监听分配的端口
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return 0, err
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}

// getFreePortInRange 从指定范围内获取空闲端口
func getFreePortInRange(min, max int) (int, error) {
	for i := 0; i < maxAttempts; i++ {
		port := rand.Intn(max-min+1) + min
		if isPortAvailable(port) {
			return port, nil
		}
	}
	return 0, fmt.Errorf("无法找到空闲端口在范围 [%d, %d]", min, max)
}

// isPortAvailable 检查指定端口是否可用
func isPortAvailable(port int) bool {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}
	listener.Close()
	return true
}
