package conf

import (
	"log"
	"os"

	"github.com/go-ini/ini"
)

// DefaultPostgreURL default url
const DefaultPostgreURL = "postgres://user:pwd@localhost/goapitest?sslmode=disable"

//DefaultPort default port
const DefaultPort = "9090"

//DefaultSwaggerBasepath swagger base path
const DefaultSwaggerBasepath = "http://localhost:9090"

// AppConf config info
var AppConf Config

// Config config struct
type Config struct {
	Dburl           string `ini:"dburl"`
	Port            string `ini:"port"`
	SwaggerBasePath string `ini:"swagger_basepath"`
}

func init() {
	log.Println("conf init")
	_, er := os.Stat("config")
	if er != nil && os.IsNotExist(er) {
		os.Mkdir("config", os.ModePerm)
	}

	ini.LooseLoad("config/app_config.ini")

	//	cfg := new(Config)
	err := ini.MapTo(&AppConf, "config/app_config.ini")

	if err != nil {
		iniFile := ini.Empty()
		conf := new(Config)
		conf.Dburl = DefaultPostgreURL
		conf.Port = DefaultPort
		conf.SwaggerBasePath = DefaultSwaggerBasepath
		ini.ReflectFrom(iniFile, conf)
		iniFile.SaveTo("config/app_config.ini")
		ini.MapTo(&AppConf, "config/app_config.ini")
	}
	log.Println("dburl:", AppConf.Dburl, ", port:", AppConf.Port)
}
