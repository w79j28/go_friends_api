// dao
package dao

import (
	"container/list"

	"github.com/w79j28/go_friends_api/entity"
	. "github.com/xormplus/xorm"
)

type UserDaoImpl struct{}

func (u *UserDaoImpl) Add(user *entity.User) int64 {
	_, err := engine.Insert(user)
	CheckError(err)
	return user.Id
}

func (u *UserDaoImpl) AddBySession(session *Session, user *entity.User) int64 {
	_, err := session.Insert(user)
	CheckError(err)
	return user.Id
}

func (u *UserDaoImpl) QueryByEmail(email string) *entity.User {
	user := new(entity.User)

	rs, _ := engine.Where("email = ?", email).Get(user)
	//CheckError(err)
	if rs {
		return user
	} else {
		return nil
	}

}

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

	var users []*entity.User = make([]*entity.User, list.Len())

	i := 0
	for e := list.Front(); e != nil; e = e.Next() {
		var uu *entity.User = e.Value.(*entity.User)

		users[i] = uu
		i++
	}
	return users
}

func (u *UserDaoImpl) Update(user *entity.User) int64 {
	//.
	affected, err := engine.Id(user.Id).Update(user)
	CheckError(err)
	return affected
}

func (u *UserDaoImpl) Delete(uid int64) int64 {

	user := new(entity.User)
	affected, err := engine.Id(uid).Delete(user)
	CheckError(err)
	return affected
}
