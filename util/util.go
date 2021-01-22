/**
 * @Author: LOFTER
 * @Description:
 * @File:  util
 * @Date: 2021/1/22 9:02 下午
 */
package util

import (
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, err
}
