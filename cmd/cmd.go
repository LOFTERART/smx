/**
 * @Author: LOFTER
 * @Description:
 * @File:  cmd
 * @Date: 2021/1/22 6:24 下午
 */
package cmd

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"smx/api"
	"smx/glo"
	"smx/middleware"
)

func Start() {

	flag.StringVar(&glo.ServerType, "t", "http", "用户名,默认为空")
	flag.StringVar(&glo.Port, "p", "80", "端口号，默认为80")
	flag.Parse()
	
	fmt.Println("serviceType:",glo.ServerType)
	fmt.Println("port:",glo.Port)

	//切换分支 启动不同的服务
	switch glo.ServerType {
	case "http":
		logfile, err := os.Create("./http.txt")
		if err != nil {
			fmt.Println("Could not create log file")
		}
		gin.SetMode(gin.DebugMode)
		gin.DefaultWriter = io.MultiWriter(logfile)
		router := gin.Default()
		router.Use(middleware.LogMiddle())
		router.GET("/download/:name", api.DownLoadFile)
		_ = router.Run(":" + glo.Port)
	case "https":
		logfile, err := os.Create("./https.txt")
		if err != nil {
			fmt.Println("Could not create log file")
		}
		gin.SetMode(gin.DebugMode)
		gin.DefaultWriter = io.MultiWriter(logfile)
		router := gin.Default()
		router.Use(middleware.LogMiddle())
		router.GET("/download/:name", api.DownLoadFile)
		_ = router.RunTLS(":"+glo.Port, "cert/server.crt", "cert/server.key")
	default:
		fmt.Println("serverType not allowed")
	}

}
