package conf

import (
	"log"
	"os"

	"github.com/go-ini/ini"
)

/**数据库连接URL**/
const DefaultPostgreUrl = "postgres://user:pwd@localhost/goapitest?sslmode=disable"
const DefaultPort = "9090"
const DefaultSwaggerBasepath = "http://localhost:9090"

var AppConf Config

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
		conf.Dburl = DefaultPostgreUrl
		conf.Port = DefaultPort
		conf.SwaggerBasePath = DefaultSwaggerBasepath
		ini.ReflectFrom(iniFile, conf)
		iniFile.SaveTo("config/app_config.ini")
		ini.MapTo(&AppConf, "config/app_config.ini")
	}
	log.Println("dburl:", AppConf.Dburl, ", port:", AppConf.Port)
}
