// user
package entity

import (
	"time"
)

type User struct {
	Id         int64
	Email      string    `xorm:"text notnull"`
	Createtime time.Time `xorm:"created"`
}

//func (h *User) TableName() string {
//	return "tbl_user"
//}
