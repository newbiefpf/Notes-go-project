package config

import (
	"github.com/BurntSushi/toml"
	"time"
)

type Config struct {
	Title     string
	App       app
	MysqlLink mysqlLink
	LogFile   logFile
}

type app struct {
	Author  string
	Info    string
	Mark    string
	Release time.Time
}
type mysqlLink struct {
	Username string
	Password string
	Host     string
	Port     int
	Dbname   string
	Timeout  string
}
type logFile struct {
	LogRouterPath string `toml:"Log_ROUTER_PATH"`
	FileName      string `toml:"File_Name"`
}

var ConfigToml Config

func ReadConfig() {
	if _, err := toml.DecodeFile("./config/config.toml", &ConfigToml); err != nil {
		panic(err)
	}
}
