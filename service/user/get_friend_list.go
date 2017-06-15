// @SubApi Friends management API [/user]
package user

import (
	"container/list"

	. "github.com/w79j28/go_friends_api/api/input"
	. "github.com/w79j28/go_friends_api/api/output"
	"github.com/w79j28/go_friends_api/dao"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @Title FriendList
// @Description   friends list for an email address.
// @Accept  json
// @Produce  json
// @Param   body     body    api.input.EmailInput     true        "EmailInput"
// @Success 200 {object} api.output.FriendListOutput "success"
// @Failure 400 {object} api.output.Output "error"
// @Router /user/friend/list [post]
// @Resource /user
func GetFriendList(c *gin.Context) {
	var json EmailInput
	result := c.BindJSON(&json)

	if result == nil {
		vFriendDao := dao.VFriendDaoImpl{}
		userList := vFriendDao.Query(json.Email)

		var output FriendListOutput
		output.Success = true

		var friends []string
		if len(userList) > 0 {
			userArray := list.New()
			for _, thisUser := range userList {
				//				fmt.Println(thisUser.User1email)
				userArray.PushBack(thisUser.User1email)
				userArray.PushBack(thisUser.User2email)
			}

			var next *list.Element
			for e := userArray.Front(); e != nil; e = next {

				if e.Value.(string) == json.Email {
					next = e.Next()
					userArray.Remove(e)
				} else {
					friends = append(friends, e.Value.(string))
					next = e.Next()
				}
			}
			output.Friends = friends
			output.Count = userArray.Len()

			//			output.Friends = []string{"www", "www2"}
			c.JSON(http.StatusOK, output)
			return
		} else {
			output.Friends = []string{}
			output.Count = 0
			c.JSON(http.StatusOK, output)
			return
		}

	}
	c.JSON(http.StatusBadRequest, FAILED)
}
