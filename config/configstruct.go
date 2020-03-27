package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

// Config 配置项结构体
type Config struct {
	Db struct {
		Mysql struct {
			DbName   string `yaml:"dbname"`
			Password string `yaml:"password"`
			Username string `yaml:"username"`
			Port     string `yaml:"port"`
			Host     string `yaml:"host"`
		}
	}
	Log struct{
		LogDirector string	`yaml:"logDirector"`
		LogAutoFile string `yaml:"logAutoFile"`
		//logWaringFile string
		LogInfoFile string `yaml:"logInfoFile"`
	}
}

// Conf 配置项
var Conf Config

// 获取文件绝对路径
func getCurrPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

var yamlFilePath string

// 获取环境变量
func (c *Config) getConfig() *Config {
	yamlFilePath = getCurrPath()
	configYaml, err := ioutil.ReadFile(yamlFilePath + "/config.yaml")
	if err != nil {
		fmt.Printf("err %v\n", err)
	}
	err = yaml.Unmarshal(configYaml, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}