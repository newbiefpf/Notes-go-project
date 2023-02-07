package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

type Config struct {
	Title string
	App   app
}

type app struct {
	Author  string
	Info    string
	Mark    string
	Release time.Time
}

var ConfigToml Config

func ReadConfig() {
	if _, err := toml.DecodeFile("./config/config.toml", &ConfigToml); err != nil {
		panic(err)
	}

	fmt.Printf("App信息：%+v\n\n", ConfigToml.App)

}
