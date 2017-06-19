package dao

import (
	"log"
)

// CheckError check error
func CheckError(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}

}
