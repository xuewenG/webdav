package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type config struct {
	Port     uint16 `yaml:"port"`
	Prefix   string `yaml:"prefix"`
	RootDir  string `yaml:"root_dir"`
	ReadOnly bool   `yaml:"read_only"`
	Users    []User `yaml:"users"`
}

var Config = &config{}

func InitConfig() error {
	configBytes, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Read config failed, %v\n", err)
		return err
	}

	err = yaml.Unmarshal(configBytes, &Config)
	if err != nil {
		log.Fatalf("Decode config failed, %v\n", err)
		return err
	}

	// 如果 Port 为空，则设置为 80
	if Config.Port == 0 {
		Config.Port = 80
	}

	// 格式化 Prefix
	Config.Prefix = fmt.Sprintf("/%s", strings.Trim(Config.Prefix, "/"))

	// 格式化 RootDir
	Config.RootDir = fmt.Sprintf("/%s", strings.Trim(Config.RootDir, "/"))

	return nil
}
