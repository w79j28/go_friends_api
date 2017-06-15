// dao
package dao

import (
	"container/list"

	"github.com/w79j28/go_friends_api/entity"
	. "github.com/xormplus/xorm"
)

type FriendDao interface {
	Add(user *entity.Friends) int64
	Query() []*entity.Friends
	Update(entity *entity.Friends) int64
	Delete(entity *entity.Friends) int64
}

type FriendDaoImpl struct{}

func (u *FriendDaoImpl) Add(entity *entity.Friends) int64 {
	id, err := engine.Insert(entity)
	CheckError(err)
	return id
}

func (u *FriendDaoImpl) AddBySession(session *Session, entity *entity.Friends) int64 {
	id, err := session.Insert(entity)
	CheckError(err)
	return id
}

func (u *FriendDaoImpl) QueryById(entity *entity.Friends) *entity.Friends {
	rs, _ := engine.Get(entity)
	//CheckError(err)
	if rs {
		return entity
	} else {
		return nil
	}
}

func (u *FriendDaoImpl) Query() []*entity.Friends {
	list := list.New()
	user := new(entity.Friends)
	rows, err := engine.Rows(user)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		user := new(entity.Friends)
		err = rows.Scan(user)
		list.PushBack(user)

	}

	var users []*entity.Friends = make([]*entity.Friends, list.Len())

	i := 0
	for e := list.Front(); e != nil; e = e.Next() {
		var uu *entity.Friends = e.Value.(*entity.Friends)

		users[i] = uu
		i++
	}
	return users
}

func (u *FriendDaoImpl) Update(user *entity.Friends) int64 {
	//.
	//	affected, err := engine.Id(user.Id).Update(user)
	//	CheckError(err)
	//	return affected
	return 0
}

func (u *FriendDaoImpl) Delete(entity *entity.Friends) int64 {

	//	user := new(entity.Friends)
	affected, err := engine.Delete(entity)
	//	affected, err := engine.Id(uid).Delete(user)
	CheckError(err)
	return affected
}
