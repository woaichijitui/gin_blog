package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"
)

// 配置文件路径
const ConfigFile = "settings.yaml"

func InitConfig() {

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

func SetYaml() error {
	//将config结构体 编码 为yaml格式
	configYamlData, err := yaml.Marshal(global.Config)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(ConfigFile, configYamlData, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
