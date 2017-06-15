package docs

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/w79j28/go_friends_api/conf"
)

var ResourceListingJson string
var ApiDescriptionsJson map[string]string

func GetSwaggerBasePathJson(str string) string {
	return strings.Replace(str, "{{basePath}}", conf.AppConf.SwaggerBasePath, -1)
}

func InitSwagger(router *gin.Engine) {
	router.StaticFS("/swagger-ui", assetFS())
	router.GET("/swagger", Swagger)

	for apiKey, _ := range ApiDescriptionsJson {
		router.GET("//"+apiKey+"/", ApiDescriptionHandler)

	}
}

func Swagger(c *gin.Context) {
	c.String(http.StatusOK, GetSwaggerBasePathJson(ResourceListingJson))
}
func ApiDescriptionHandler(c *gin.Context) {
	apiKey := strings.Trim(c.Request.RequestURI, "/")

	if json, ok := ApiDescriptionsJson[apiKey]; ok {
		json = GetSwaggerBasePathJson(json)
		_, e := template.New("validate").Parse(json)
		if e != nil {
			c.Status(http.StatusNotFound)
			return
		}
		//
		c.String(http.StatusOK, (json))
	} else {
		c.Status(http.StatusNotFound)
	}
}
