package config

import (
	"errors"
	"gopkg.in/ini.v1"
	"log"
)
type Config struct {
	AppName   string `ini:"app_name"`
	LogLevel  string `ini:"log_level"`

	MySQL     MySQLConfig `ini:"mysql"`
	Redis     RedisConfig `ini:"redis"`
}

type MySQLConfig struct {
	IP        string `ini:"ip"`
	Port      int `ini:"port"`
	User      string `ini:"user"`
	Password  string `ini:"password"`
	Database  string `ini:"database"`
}

type RedisConfig struct {
	IP      string `ini:"ip"`
	Port    int `ini:"port"`
}

func Init(cfgName string) (*ini.File, error) {
	load, err := ini.Load(cfgName)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return load, nil
}

// ParseConfig 整体映射到对应结构体
func ParseConfig(cfg *ini.File, v interface{}) error {
	if cfg == nil {
		return errors.New("cfg is nil")
	}

	err := cfg.MapTo(v)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// ParseSection 将section分区映射到对应结构体
func ParseSection(cfg *ini.File, section string, v interface{}) error {
	if cfg == nil {
		return errors.New("cfg is nil")
	}

	sec, err := cfg.GetSection(section)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = sec.MapTo(v)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}