/**
 * @Author: LOFTER
 * @Description:
 * @File:  main
 * @Date: 2021/1/22 10:20 下午
 */
package main

import (
	"smx/cmd"
	"smx/model"
)

func init() {
	model.Initialized()
}

func main() {
	cmd.Start()
}
