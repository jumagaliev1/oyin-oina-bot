package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"sync"
)

type Config struct {
	Bot struct {
		Token string `yaml:"token" env:"BOT_TOKEN"`
	} `yaml:"bot"`
}

var configPath string
var instance *Config
var once sync.Once

const (
	EnvConfigPathName  = "CONFIG-PATH"
	FlagConfigPathName = "config"
)

func GetConfig() *Config {
	once.Do(func() {
		flag.StringVar(&configPath, FlagConfigPathName, "", "this is app config file")
		flag.Parse()

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}
		if configPath == "" {
			log.Fatal("config path is required")
		}

		instance = &Config{}
		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			var headerText = "Bot config"
			helpText, _ := cleanenv.GetDescription(instance, &headerText)
			log.Print(helpText)
			log.Fatal(err)
		}
	})
	return instance
}
