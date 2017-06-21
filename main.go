// @APIVersion 1.0.0
// @APITitle Swagger Friends Management API
// @APIDescription Swagger Friends Management API
// @BasePath {{basePath}}
// @Contact jingzhu.wang@chinasofti.com
// @TermsOfServiceUrl http://www.chinasofti.com
// @License None
// @LicenseUrl #
package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/w79j28/go_friends_api/conf"
	"github.com/w79j28/go_friends_api/service/user"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/w79j28/go_friends_api/swagger.docs"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.Use(corsHandler)

	userRouter := router.Group("/user")
	{

		userRouter.POST("/friends", user.AddFriend)
		userRouter.POST("/friend/list", user.GetFriendList)
		userRouter.POST("/friends/common", user.GetCommonFriends)
		userRouter.POST("/friend/subscribe", user.SubscribeFriend)
		userRouter.POST("/friend/block", user.BlockFriend)
		userRouter.POST("/friends/sender", user.SenderFriends)

	}

	//
	docs.InitSwagger(router)
	if conf.AppConf.Port == "cloud" {
		router.Run(":" + os.Getenv("PORT"))
	} else {
		router.Run(":" + conf.AppConf.Port)
	}
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
