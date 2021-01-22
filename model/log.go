/**
 * @Author: LOFTER
 * @Description:
 * @File:  log
 * @Date: 2021/1/22 9:40 下午
 */
package model

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Time      string `json:"time"`
}
