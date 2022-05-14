package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Yaml struct {
	MySQL struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	}
	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	JwtKey string
}

var (
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string

	RedisHost string
	RedisPort string

	JwtKey string
)

func init() {
	var conf Yaml
	file, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
		return
	}
	if err = yaml.Unmarshal(file, &conf); err != nil {
		fmt.Println("配置文件解析错误：", err)
		return
	}
	LoadDB(&conf)
	LoadRedis(&conf)
	LoadJwt(&conf)
}
func LoadDB(conf *Yaml) {
	DbUser = conf.MySQL.User
	DbPassword = conf.MySQL.Password
	DbHost = conf.MySQL.Host
	DbPort = conf.MySQL.Port
	DbName = conf.MySQL.Name
}
func LoadRedis(conf *Yaml) {
	RedisHost = conf.Redis.Host
	RedisPort = conf.Redis.Port
}
func LoadJwt(conf *Yaml) {
	JwtKey = conf.JwtKey
}
