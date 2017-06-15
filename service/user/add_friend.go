// @SubApi Friends management API [/user]
package user

import (
	. "github.com/w79j28/go_friends_api/api/input"
	. "github.com/w79j28/go_friends_api/api/output"
	"github.com/w79j28/go_friends_api/dao"

	"net/http"

	"github.com/w79j28/go_friends_api/entity"

	"github.com/w79j28/go_friends_api/util"

	"github.com/gin-gonic/gin"
)

// @Title AddFriend
// @Description  create a friend connection between two email addresses.
// @Accept  json
// @Produce  json
// @Param   body     body    api.input.FriendInput     true        "FriendInput"
// @Success 200 {object} api.output.Output "success"
// @Failure 400 {object} api.output.Output "error"
// @Router /user/friends [post]
// @Resource /user
func AddFriend(c *gin.Context) {

	//name := c.Param("name")
	//c.String(http.StatusOK, "Hello %s", name)
	//c.JSON(200, gin.H{"name": name})
	//	var userInfo LoginForm
	//	userInfo.User = name
	//	userInfo.Password = name + "password"
	//	c.JSON(200, userInfo)
	//	fmt.Println("add friend....")
	var json FriendInput
	result := c.BindJSON(&json)
	if result == nil {
		intLen := len(json.Friends)
		if intLen != 2 {
			c.JSON(http.StatusBadRequest, FAILED)
			return
		}
		if json.Friends[0] == json.Friends[1] {
			c.JSON(http.StatusBadRequest, FAILED)
			return
		}

		userDao := dao.UserDaoImpl{}
		var userId []int64
		var notExistUserCount int
		session := dao.NewSessionBegin()
		defer dao.SessionDeferFunc(session, func() {
			c.JSON(http.StatusBadRequest, FAILED)
		})

		for _, thisEmail := range json.Friends {
			existUser := userDao.QueryByEmail(thisEmail)
			if existUser == nil {
				user := new(entity.User)
				user.Email = thisEmail
				userDao.AddBySession(session, user)
				userId = append(userId, user.Id)
				notExistUserCount++
			} else {
				//已存在
				userId = append(userId, existUser.Id)

			}
		}
		vPermissionDao := dao.VPermissionDaoImpl{}
		permissionList := vPermissionDao.QueryByTargetEmailAndStatus(BLOCK, json.Friends[0], json.Friends[1])
		user0HashSet := util.NewHashSet()
		user1HashSet := util.NewHashSet()
		for _, thisPermission := range permissionList {
			if thisPermission.Targetemail == json.Friends[0] {
				user0HashSet.Add(thisPermission.Requestoremail)
			} else if thisPermission.Targetemail == json.Friends[1] {
				user1HashSet.Add(thisPermission.Requestoremail)
			}

		}
		if user0HashSet.Contains(json.Friends[1]) || user1HashSet.Contains(json.Friends[0]) {
			c.JSON(http.StatusBadRequest, FAILED)
			return
		}

		util.QuickSort(userId)
		friend := new(entity.Friends)
		friend.Userid1 = userId[0]
		friend.Userid2 = userId[1]
		friendDao := dao.FriendDaoImpl{}
		if friendDao.QueryById(friend) == nil {
			friendDao.AddBySession(session, friend)
		}

		c.JSON(http.StatusCreated, SUCCESS)
	} else {
		//		fmt.Println("failed:", result)
		c.JSON(http.StatusBadRequest, FAILED)
	}
}
