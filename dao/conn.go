// conn
package dao

import (
	//	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/w79j28/go_friends_api/conf"
	"github.com/w79j28/go_friends_api/entity"

	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

/**数据库连接URL**/
//const postgreUrl = conf.PostgreUrl

/****/
var engine *xorm.Engine
var err error

func init() {
	log.Println("dao init")
	DbInit()
	SyncTable()
}

func NewSessionBegin() *xorm.Session {
	session := engine.NewSession()
	session.Begin()
	return session
}

func SessionDeferFunc(session *xorm.Session, failedFunc func()) {
	if info := recover(); info != nil {
		session.Rollback()
		session.Close()
		failedFunc()
		return
	} else {
		session.Commit()
		session.Close()
	}

}

/**
 * 初始化
 *
 */
func DbInit() {
	engine, err = xorm.NewPostgreSQL(conf.AppConf.Dburl)
	CheckError(err)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "api_")
	engine.SetTableMapper(tbMapper)

	engine.ShowSQL(true)
}

/**
 *同步表
 *
 */
func SyncTable() {
	err := engine.Sync2(new(entity.User), new(entity.Friends), new(entity.Permission))

	CheckError(err)
	engine.ImportFile("./conf/db.sql")

}
