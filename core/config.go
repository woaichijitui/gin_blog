package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"
)

func InitConfig() {

	//配置文件路径
	const ConfigFile = "settings.yaml"

	//ioutil.ReadFile 将文件读取到[]byte
	configFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Sprintf("配置文件读取错误：%v", err))
	}

	yamlConfig := config.Config{}
	err = yaml.Unmarshal(configFile, &yamlConfig)
	if err != nil {
		log.Fatalf("config Init Unmarshal : %v", err)
	}
	log.Println("config yamlFile load init success.")
	global.Config = &yamlConfig

}
