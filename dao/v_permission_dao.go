package dao

import (
	"github.com/w79j28/go_friends_api/entity"
)

// VPermissionDaoImpl api_permission view
type VPermissionDaoImpl struct{}

// QueryByTargetEmail queryByTargetEmail
func (u *VPermissionDaoImpl) QueryByTargetEmail(email ...string) []entity.VPermission {
	var users []entity.VPermission
	err := engine.Table("v_permission").In("targetemail", email).Asc("status").Find(&users)
	CheckError(err)
	return users
}

// QueryByTargetEmailAndStatus queryByTargetEmailAndStatus
func (u *VPermissionDaoImpl) QueryByTargetEmailAndStatus(status int, email ...string) []entity.VPermission {
	var users []entity.VPermission
	err := engine.Table("v_permission").Where("status = ?", status).In("targetemail", email).Asc("status").Find(&users)
	CheckError(err)
	return users
}
