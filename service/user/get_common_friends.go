// @SubApi Friends management API [/user]
package user

import (
	"container/list"
	"fmt"
	"net/http"

	"github.com/w79j28/go_friends_api/api/input"
	"github.com/w79j28/go_friends_api/api/output"
	"github.com/w79j28/go_friends_api/dao"

	"github.com/gin-gonic/gin"
)

// @Title GetCommonFriends
// @Description  common friends list between two email addresses.
// @Accept  json
// @Produce  json
// @Param   body     body    api.input.FriendInput     true        "FriendInput"
// @Success 200 {object} api.output.FriendListOutput "success"
// @Failure 400 {object} api.output.Output "error"
// @Router /user/friends/common [post]
// @Resource /user
func GetCommonFriends(c *gin.Context) {
	var json input.FriendInput
	result := c.BindJSON(&json)
	if result == nil {
		intLen := len(json.Friends)
		if intLen != 2 {
			c.JSON(http.StatusBadRequest, output.FAILED)
			return
		}

		defer func() {
			if info := recover(); info != nil {
				c.JSON(http.StatusBadRequest, output.FAILED)
				return
			}
		}()
		userDao := dao.VFriendDaoImpl{}
		userList := userDao.QueryByEmail(json.Friends[0], json.Friends[1])
		user1List, user2List := list.New(), list.New()
		for _, thisUser := range userList {
			// user 1 list
			if thisUser.User1email == json.Friends[0] {
				user1List.PushBack(thisUser.User2email)
			} else if thisUser.User2email == json.Friends[0] {
				user1List.PushBack(thisUser.User1email)
			}

			//user 2 list
			if thisUser.User1email == json.Friends[1] {
				user2List.PushBack(thisUser.User2email)
			} else if thisUser.User2email == json.Friends[1] {
				user2List.PushBack(thisUser.User1email)
			}

		}
		var commons = []string{}
		for e := user1List.Front(); e != nil; e = e.Next() {
			for e2 := user2List.Front(); e2 != nil; e2 = e2.Next() {
				if e.Value == e2.Value {
					commons = append(commons, e.Value.(string))
				}
			}
		}

		var output output.FriendListOutput
		output.Success = true
		output.Count = len(commons)
		output.Friends = commons
		c.JSON(http.StatusOK, output)
	} else {
		fmt.Println("failed:", result)
		c.JSON(http.StatusBadRequest, output.FAILED)
	}
}
