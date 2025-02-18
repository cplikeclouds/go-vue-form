package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"server/config"
	"server/global"

	"io/ioutil"
	"log"
)

// InitConf 读取yaml文件配置
func InitConf() {

	const ConfigFile = "config.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get YAML config file error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Unmarshal YAML config file error:%s", err)
	}
	log.Println("Config yaml load Init success")
	global.Config = c
	
}
