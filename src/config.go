package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Hostname string `yaml:"Hostname"`
	Port     string `yaml:"Port"`
}

var gCfg config

// 从配置文件读取最新的配置
func updateConfig() (config, error) {
	gCfg = config{}
	bs, err := ioutil.ReadFile("config/server.yaml")
	if err != nil {
		return gCfg, err
	}
	if err := yaml.Unmarshal(bs, &gCfg); err != nil {
		return gCfg, err
	}
	return gCfg, nil
}

// 获取服务器的绑定地址
func getAddr() (string, error) {
	cfg, err := updateConfig()
	if err != nil {
		return "", err
	}
	addr := cfg.Hostname + ":" + cfg.Port
	if cfg.Hostname == "" && cfg.Port == "" {
		return "", fmt.Errorf("请指定正确的绑定地址")
	}
	return addr, nil
}

// 初始化所有的文件夹指定
func initDirs() error {
	return nil
}
