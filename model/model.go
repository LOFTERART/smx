/**
 * @Author: LOFTER
 * @Description:
 * @File:  model
 * @Date: 2021/1/22 9:51 下午
 */
package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

//Initialized  连接mysql
func Initialized() {

	//var err error
	
	DB,_= gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	
	_ = DB.AutoMigrate(&Log{})
	
}
