package dao

import (
	"github.com/w79j28/go_friends_api/entity"
)

// VFriendDaoImpl  api_friends view
type VFriendDaoImpl struct{}

// Query query
func (u *VFriendDaoImpl) Query(email string) []entity.VFriends {
	var users []entity.VFriends
	err := engine.Table("v_friends").Where("user1email = ?", email).Or("user2email = ?", email).Find(&users)
	CheckError(err)
	return users
}

// QueryByEmail queryByEmail
func (u *VFriendDaoImpl) QueryByEmail(email1, email2 string) []entity.VFriends {
	var users []entity.VFriends
	err := engine.Table("v_friends").Where("user1email = ?", email1).Or("user1email = ?", email2).Or("user2email = ?", email1).Or("user2email = ?", email2).Find(&users)
	CheckError(err)
	return users
}
