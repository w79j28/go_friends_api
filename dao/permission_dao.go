// dao
package dao

import (
	"github.com/w79j28/go_friends_api/entity"
	//	"fmt"
	. "github.com/xormplus/xorm"
)

type PermissionDaoImpl struct{}

func (u *PermissionDaoImpl) Add(entity *entity.Permission) int64 {
	id, err := engine.Insert(entity)
	CheckError(err)
	return id
}

func (u *PermissionDaoImpl) AddBySession(session *Session, entity *entity.Permission) int64 {
	id, err := session.Insert(entity)
	CheckError(err)
	return id
}

func (u *PermissionDaoImpl) QueryByEmail(email string) *entity.Permission {
	user := new(entity.Permission)

	rs, _ := engine.Where("email = ?", email).Get(user)
	//CheckError(err)
	if rs {
		return user
	} else {
		return nil
	}

}

func (u *PermissionDaoImpl) Query(entity *entity.Permission) *entity.Permission {
	rs, _ := engine.Get(entity)
	if rs {
		return entity
	} else {
		return nil
	}
}

//func (u *PermissionDaoImpl) Query() []*entity.Permission {
//	list := list.New()
//	user := new(entity.Permission)
//	rows, err := engine.Rows(user)
//	CheckError(err)
//	defer rows.Close()
//	for rows.Next() {
//		user := new(entity.User)
//		err = rows.Scan(user)
//		list.PushBack(user)

//	}

//	var users []*entity.User = make([]*entity.Permission, list.Len())

//	i := 0
//	for e := list.Front(); e != nil; e = e.Next() {
//		var uu *entity.Permission = e.Value.(*entity.Permission)

//		users[i] = uu
//		i++
//	}
//	return users
//}

//func (u *PermissionDaoImpl) Update(user *entity.Permission) int64 {
//	//.
//	affected, err := engine.Id(user.Id).Update(user)
//	CheckError(err)
//	return affected
//}

//func (u *PermissionDaoImpl) Delete(uid int64) int64 {

//	user := new(entity.Permission)
//	affected, err := engine.Id(uid).Delete(user)
//	CheckError(err)
//	return affected
//}
