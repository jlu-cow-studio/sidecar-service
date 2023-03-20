package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jlu-cow-studio/common/discovery"
)

const (
	ENV_SERVICE_NAME   = "ENV_SERVICE_NAME"
	ENV_ADDRESS        = "ENV_ADDRESS"
	ENV_PORT           = "EVN_PORT"
	ENV_SIDECAR_PORT   = "ENV_SIDECAR_PORT"
	INNER_SIDECAR_PORT = "8081"
)

func main() {
	discovery.Init()

	serviceName := os.Args[1]
	serviceAddress := os.Args[2]
	servicePort := os.Args[3]
	sidecarPort := os.Args[4]

	if serviceName == "" {
		panic(errors.New("empty serviceName"))
	}

	if serviceAddress == "" {
		panic(errors.New("empty serviceAddress"))
	}

	if servicePort == "" {
		panic(errors.New("empty servicePort"))
	}

	if sidecarPort == "" {
		panic(errors.New("empty sidecarPort"))
	}

	port, err := strconv.Atoi(servicePort)
	if err != nil {
		panic(err)
	}

	// 注册服务
	if err = discovery.Register(serviceName, serviceAddress, fmt.Sprintf("http://%s:%s/health", serviceAddress, sidecarPort), port); err != nil {
		panic(err)
	}

	// 创建gin引擎
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		fmt.Fprintln(c.Writer, "OK")
	})

	// 启动HTTP服务，默认监听地址为:8080
	router.Run(fmt.Sprintf("0.0.0.0:%s", INNER_SIDECAR_PORT))
}
