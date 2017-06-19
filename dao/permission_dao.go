package dao

import (
	"github.com/w79j28/go_friends_api/entity"
	"github.com/xormplus/xorm"
)

// PermissionDaoImpl PermissionDao struct
type PermissionDaoImpl struct{}

// Add add
func (u *PermissionDaoImpl) Add(entity *entity.Permission) int64 {
	id, err := engine.Insert(entity)
	CheckError(err)
	return id
}

// AddBySession addBySession
func (u *PermissionDaoImpl) AddBySession(session *xorm.Session, entity *entity.Permission) int64 {
	id, err := session.Insert(entity)
	CheckError(err)
	return id
}

// QueryByEmail queryByEmail
func (u *PermissionDaoImpl) QueryByEmail(email string) *entity.Permission {
	user := new(entity.Permission)

	rs, _ := engine.Where("email = ?", email).Get(user)
	//CheckError(err)
	if rs {
		return user
	}
	return nil

}

// Query query
func (u *PermissionDaoImpl) Query(entity *entity.Permission) *entity.Permission {
	rs, _ := engine.Get(entity)
	if rs {
		return entity
	}
	return nil

}
