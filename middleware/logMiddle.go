/**
 * @Author: LOFTER
 * @Description:
 * @File:  logMiddle
 * @Date: 2021/1/22 9:47 下午
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"smx/model"
	"time"
)

type LogInfo struct {
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Time      string `json:"time"`
}

func LogMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		info:=model.Log{
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Time:      time.Now().Format("2006-01-02 15:04:05"),
		}
		model.DB.Create(&info)
		c.Next()
	}
}
