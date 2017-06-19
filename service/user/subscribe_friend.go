// @SubApi Friends management API [/user]
package user

import (
	"net/http"

	"github.com/w79j28/go_friends_api/api/input"
	"github.com/w79j28/go_friends_api/api/output"
	"github.com/w79j28/go_friends_api/dao"
	"github.com/w79j28/go_friends_api/entity"

	"github.com/gin-gonic/gin"
)

const SUBSCRIBE int = 1

// @Title SubscribeFriend
// @Description  subscribing to updates" is NOT equivalent to "adding a friend connection.
// @Accept  json
// @Produce  json
// @Param   body     body    api.input.SubscribeInput     true        "SubscribeInput"
// @Success 200 {object} api.output.Output "success"
// @Failure 400 {object} api.output.Output "error"
// @Router /user/friend/subscribe [post]
// @Resource /user
func SubscribeFriend(c *gin.Context) {
	var json input.SubscribeInput
	result := c.BindJSON(&json)
	if result == nil {
		if json.Requestor == json.Target {
			c.JSON(http.StatusBadRequest, output.FAILED)
			return
		}
		session := dao.NewSessionBegin()
		defer dao.SessionDeferFunc(session, func() {
			c.JSON(http.StatusBadRequest, output.FAILED)
		})
		userDao := dao.UserDaoImpl{}
		requestorUser := userDao.QueryByEmail(json.Requestor)
		if requestorUser == nil {
			// Requestor 不存在
			user := new(entity.User)
			user.Email = json.Requestor
			userDao.AddBySession(session, user)
			requestorUser = user
		}
		targetUser := userDao.QueryByEmail(json.Target)
		if targetUser == nil {
			// Target 不存在
			user := new(entity.User)
			user.Email = json.Target
			userDao.AddBySession(session, user)
			targetUser = user
		}

		entity := entity.Permission{}
		entity.Requestor = requestorUser.Id
		entity.Target = targetUser.Id
		entity.Status = SUBSCRIBE

		permissionDao := dao.PermissionDaoImpl{}
		thisPermission := permissionDao.Query(&entity)
		if thisPermission == nil {
			permissionDao.AddBySession(session, &entity)
		}
		c.JSON(http.StatusCreated, output.SUCCESS)
	} else {
		c.JSON(http.StatusBadRequest, output.FAILED)
	}
}
