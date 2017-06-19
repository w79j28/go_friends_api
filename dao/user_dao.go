package dao

import (
	"container/list"

	"github.com/w79j28/go_friends_api/entity"
	"github.com/xormplus/xorm"
)

// UserDaoImpl userDao struct
type UserDaoImpl struct{}

//Add add
func (u *UserDaoImpl) Add(user *entity.User) int64 {
	_, err := engine.Insert(user)
	CheckError(err)
	return user.Id
}

// AddBySession addBySession
func (u *UserDaoImpl) AddBySession(session *xorm.Session, user *entity.User) int64 {
	_, err := session.Insert(user)
	CheckError(err)
	return user.Id
}

// QueryByEmail queryByEmail
func (u *UserDaoImpl) QueryByEmail(email string) *entity.User {
	user := new(entity.User)

	rs, _ := engine.Where("email = ?", email).Get(user)
	//CheckError(err)
	if rs {
		return user
	}
	return nil

}

// Query query
func (u *UserDaoImpl) Query() []*entity.User {
	list := list.New()
	user := new(entity.User)
	rows, err := engine.Rows(user)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		user := new(entity.User)
		err = rows.Scan(user)
		list.PushBack(user)

	}

	var users = make([]*entity.User, list.Len())

	i := 0
	for e := list.Front(); e != nil; e = e.Next() {
		var uu = e.Value.(*entity.User)
		users[i] = uu
		i++
	}
	return users
}

// Update update
func (u *UserDaoImpl) Update(user *entity.User) int64 {
	//.
	affected, err := engine.Id(user.Id).Update(user)
	CheckError(err)
	return affected
}

// Delete delete
func (u *UserDaoImpl) Delete(uid int64) int64 {

	user := new(entity.User)
	affected, err := engine.Id(uid).Delete(user)
	CheckError(err)
	return affected
}
