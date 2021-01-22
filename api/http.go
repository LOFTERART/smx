/**
 * @Author: LOFTER
 * @Description:
 * @File:  http
 * @Date: 2021/1/22 8:12 下午
 */
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"smx/util"
	"time"
)

//hello
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"ip":         c.ClientIP(),
		"user_agent": c.Request.UserAgent(),
		"time":       time.Now(),
	})
}

//download txt

func DownLoadFile(c *gin.Context) {
	name := c.Param("name")
	switch len(name) {
	case 3,4,5,6:
		if _, err := util.PathExists(name+".txt"); err != nil {
			c.JSON(200, "File Not Exit,Create a *.txt file in the root directory ")
		} else {
			c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name+".txt"))
			c.Writer.Header().Add("Content-Type", "application/octet-stream")
			c.File(name+".txt")
		}
	default:
		c.JSON(200, "not allowed download")
	}

}
