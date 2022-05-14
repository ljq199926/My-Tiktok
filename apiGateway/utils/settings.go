package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Yaml struct {
	JwtKey string
}

var (
	JwtKey string
)

func init() {
	var conf Yaml
	file, err := ioutil.ReadFile("deploy/conf.yaml")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
		return
	}
	if err = yaml.Unmarshal(file, &conf); err != nil {
		fmt.Println("配置文件解析错误：", err)
		return
	}
	LoadJwt(&conf)
}
func LoadJwt(conf *Yaml) {
	JwtKey = conf.JwtKey
}
