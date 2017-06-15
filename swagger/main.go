package main

import (
	"github.com/yvasiyarov/swagger/generator"
)

func main() {
	var apiPackage string = "github.com\\w79j28\\go_friends_api\\service"
	var basePackage = "github.com/w79j28/go_friends_api"
	var mainApiFile string = "github.com/w79j28/go_friends_api/main.go"
	var outputFormat string = "go"
	var outputSpec string = "../swagger.docs"
	var controllerClass string = ""
	var ignore string = "^$"
	var contentsTable bool = true
	var models bool = false
	var vendoringPath string = "github.com/w79j28/go_friends_api/vendor"

	params := generator.Params{
		ApiPackage:      apiPackage,
		BasePackage:     basePackage,
		MainApiFile:     mainApiFile,
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
