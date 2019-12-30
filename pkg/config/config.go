// username: vonhng
// create_time: 2019/12/29 - 22:21
// mail: vonhng.feng@gmail.com
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	DB  MongoDB    `yaml:"mongodb"`
	Web WebService `yaml:"web"`
}

type MongoDB struct {
	Host     string `json:"host";yaml:"host"`
	Port     string `json:"port";yaml:"port"`
	Username string `json:"username";yaml:"username"`
	Password string `json:"password";yaml:"password"`
	DBName   string `json:"db_name";yaml:"db_name"`
}

type WebService struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

var Setting = &Config{}

func ParseConfig() {
	data, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil {
		log.Fatal("missing config.yaml")
	}
	if err = yaml.Unmarshal(data, &Setting); err != nil {
		log.Fatalf("Unmarshal test.yaml err: %s", err)
	}
}
