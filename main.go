// @APIVersion 1.0.0
// @APITitle Swagger Friends Management API
// @APIDescription Swagger Friends Management API
// @BasePath {{basePath}}
// @Contact jingzhu.wang@chinasofti.com
// @TermsOfServiceUrl http://www.chinasofit.com
// @License None
// @LicenseUrl #
package main

import (
	"fmt"
	"strings"
	//	"fmt"
	"net/http"
	//	"strings"

	"github.com/w79j28/go_friends_api/conf"
	. "github.com/w79j28/go_friends_api/service/user"

	. "github.com/w79j28/go_friends_api/swagger.docs"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.Use(corsHandler)

	user := router.Group("/user")
	{

		user.POST("/friends", AddFriend)
		user.POST("/friend/list", GetFriendList)
		user.POST("/friends/common", GetCommonFriends)
		user.POST("/friend/subscribe", SubscribeFriend)
		user.POST("/friend/block", BlockFriend)
		user.POST("/friends/sender", SenderFriends)

	}

	//
	InitSwagger(router)
	router.Run(":" + conf.AppConf.Port)

}

func corsHandler(c *gin.Context) {
	fmt.Println("RequestURI :")
	if strings.HasPrefix(c.Request.RequestURI, "/user") {
		c.Header("Content-Type", "application/json")
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "origin, content-type, accept, authorization, Pragma, Cache-control, Expires")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	c.Header("Access-Control-Max-Age", "1209600")
}
