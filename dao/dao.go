package dao

import (
	"log"
)

/**
 * 错误检测
 *
 */
func CheckError(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}

}
