package main

import (
	"github.com/yvasiyarov/swagger/generator"
)

func main() {
	var apiPackage = "github.com\\w79j28\\go_friends_api\\service"
	var basePackage = "github.com/w79j28/go_friends_api"
	var mainAPIFile = "github.com/w79j28/go_friends_api/main.go"
	var outputFormat = "go"
	var outputSpec = "../swagger.docs"
	var controllerClass = ""
	var ignore = "^$"
	var contentsTable = true
	var models = false
	var vendoringPath = "github.com/w79j28/go_friends_api/vendor"

	params := generator.Params{
		ApiPackage:      apiPackage,
		BasePackage:     basePackage,
		MainApiFile:     mainAPIFile,
		OutputFormat:    outputFormat,
		OutputSpec:      outputSpec,
		ControllerClass: controllerClass,
		Ignore:          ignore,
		ContentsTable:   contentsTable,
		Models:          models,
		VendoringPath:   vendoringPath,
	}

	err := generator.Run(params)
	if err != nil {
		//		log.Fatal(err.Error())
	}
}
