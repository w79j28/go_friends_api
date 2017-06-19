// @SubApi Friends management API [/user]
package user

import (
	"net/http"

	"github.com/w79j28/go_friends_api/dao"
	"github.com/w79j28/go_friends_api/util"

	"github.com/w79j28/go_friends_api/api/input"
	"github.com/w79j28/go_friends_api/api/output"

	"github.com/gin-gonic/gin"
)

// @Title SenderFriends
// @Description  all email addresses that can receive updates from an email address.
// @Accept  json
// @Produce  json
// @Param   body     body    api.input.SenderInput     true    "SenderInput"
// @Success 200 {object} api.output.RecipientsOutput "success"
// @Failure 400 {object} api.output.Output "error"
// @Router /user/friends/sender [post]
// @Resource /user
func SenderFriends(c *gin.Context) {
	var json input.SenderInput

	result := c.BindJSON(&json)
	if result == nil {
		defer func() {
			if info := recover(); info != nil {
				c.JSON(http.StatusBadRequest, output.FAILED)
				return
			}
		}()
		vFriendDao := dao.VFriendDaoImpl{}
		vPermissionDao := dao.VPermissionDaoImpl{}

		// Eligibility for receiving updates from i.e. "john@example.com":
		//•  has not blocked updates from "john@example.com", and
		//•  at least one of the following:
		//   o  has a friend connection with "john@example.com"
		//   o  has subscribed to updates from "john@example.com"
		//   o  has been @mentioned in the update
		friendList := vFriendDao.Query(json.Sender)
		permissionList := vPermissionDao.QueryByTargetEmail(json.Sender)

		userHashSet := util.NewHashSet()
		for _, thisFriend := range friendList {
			userHashSet.Add(thisFriend.User1email)
			userHashSet.Add(thisFriend.User2email)
		}
		for _, thisPermission := range permissionList {
			if thisPermission.Status == 1 {
				userHashSet.Add(thisPermission.Requestoremail)

			} else if thisPermission.Status == 2 {
				userHashSet.Remove(thisPermission.Requestoremail)
			}
		}
		var recipientList = []string{}
		for _, element := range userHashSet.Elements() {
			if element.(string) != json.Sender {
				recipientList = append(recipientList, element.(string))
			}
		}

		output := output.RecipientsOutput{}
		output.Success = true
		output.Recipients = recipientList
		c.JSON(http.StatusOK, output)
	} else {

		c.JSON(http.StatusBadRequest, output.FAILED)
	}
}
