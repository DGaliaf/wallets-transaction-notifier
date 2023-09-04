package cfg

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Bot struct {
		Token string `yaml:"token"`
	} `yaml:"bot"`
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"db"`
}

var config *Config
var once sync.Once

func GetConfig(path string) *Config {
	config = &Config{}

	once.Do(func() {
		if err := cleanenv.ReadConfig(path, config); err != nil {
			help, _ := cleanenv.GetDescription(config, nil)

			log.Println(help)
		}
	})

	return config
}
