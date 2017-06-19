package docs

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/w79j28/go_friends_api/conf"
)

// ResourceListingJSON Resource listing json
var ResourceListingJSON string

// APIDescriptionsJSON Api Descriptions Json
var APIDescriptionsJSON map[string]string

// GetSwaggerBasePathJson Get SwaggerBasePath Json
func GetSwaggerBasePathJson(str string) string {
	return strings.Replace(str, "{{basePath}}", conf.AppConf.SwaggerBasePath, -1)
}

// InitSwagger Swagger url
func InitSwagger(router *gin.Engine) {
	router.StaticFS("/swagger-ui", assetFS())
	router.GET("/swagger", Swagger)

	for apiKey, _ := range APIDescriptionsJSON {
		router.GET("//"+apiKey+"/", ApiDescriptionHandler)

	}
}

// Swagger SwaggerJSON
func Swagger(c *gin.Context) {
	c.String(http.StatusOK, GetSwaggerBasePathJson(ResourceListingJSON))
}

// ApiDescriptionHandler swagger url  Handler
func ApiDescriptionHandler(c *gin.Context) {
	apiKey := strings.Trim(c.Request.RequestURI, "/")

	if json, ok := APIDescriptionsJSON[apiKey]; ok {
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
