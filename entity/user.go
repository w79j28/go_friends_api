package entity

import (
	"time"
)

// User table api_user
type User struct {
	Id         int64
	Email      string    `xorm:"text notnull"`
	Createtime time.Time `xorm:"created"`
}
