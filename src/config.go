package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type serveConfig struct {
	Hostname string `yaml:"Hostname"`
	Port     string `yaml:"Port"`
}
type dirConfig struct {
	ResourceDir string `yaml:"ResourceDir"`
}
type ResTypeConfig struct {
	Actions map[string]bool `yaml:"Actions"`
	Styles  map[string]bool `yaml:"Styles"`
	Images  map[string]bool `yaml:"Images"`
	Fonts   map[string]bool `yaml:"Fonts"`
	Pages   map[string]bool `yaml:"Pages"`
}

// 从配置文件读取最新的配置
func updateServeConfig() (serveConfig, error) {
	sCfg := serveConfig{}
	bs, err := ioutil.ReadFile("config/server.yaml")
	if err != nil {
		return sCfg, err
	}
	if err := yaml.Unmarshal(bs, &sCfg); err != nil {
		return sCfg, err
	}
	return sCfg, nil
}

func updateDirConfig() (dirConfig, error) {
	dCfg := dirConfig{}
	bs, err := ioutil.ReadFile("config/dir.yaml")
	if err != nil {
		return dCfg, err
	}
	if err := yaml.Unmarshal(bs, &dCfg); err != nil {
		return dCfg, err
	}
	return dCfg, nil
}

func updateResTypeConfig() (ResTypeConfig, error) {
	rCfg := ResTypeConfig{}
	bs, err := ioutil.ReadFile("config/resType.yaml")
	if err != nil {
		return rCfg, err
	}
	if err := yaml.Unmarshal(bs, &rCfg); err != nil {
		return rCfg, err
	}
	return rCfg, nil
}

func getAddr() (string, error) {
	cfg, err := updateServeConfig()
	if err != nil {
		return "", err
	}
	addr := cfg.Hostname + ":" + cfg.Port
	if cfg.Hostname == "" && cfg.Port == "" {
		return "", fmt.Errorf("请指定正确的绑定地址")
	}
	return addr, nil
}

func getResourceDir() (string, error) {
	cfg, err := updateDirConfig()
	if err != nil {
		return "", err
	}
	return cfg.ResourceDir, nil
}

func getResType(suffix string) (uint, error) {
	cfg, err := updateResTypeConfig()
	if err != nil {
		return 0, err
	}
	f := suffix[1:]
	if cfg.Actions[f] {
		return ResType_Actions, nil
	}
	if cfg.Styles[f] {
		return ResType_Styles, nil
	}
	if cfg.Images[f] {
		return ResType_Images, nil
	}
	if cfg.Fonts[f] {
		return ResType_Images, nil
	}
	if cfg.Pages[f] {
		return ResType_Images, nil
	}
	return 0, nil
}
