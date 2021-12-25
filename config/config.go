package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var cfg Config
var doOnce sync.Once

type Config struct {
	Application struct {
		Port int `mapstructure:"PORT"`
		Log  struct {
			Path string `mapstructure:"PATH"`
		}
		Key struct {
			Default string `mapstructure:"DEFAULT"`
			Rsa     struct {
				Public  string `mapstructure:"PUBLIC"`
				Private string `mapstructure:"PRIVATE"`
			}
		}
	} `mapstructure:"APPLICATION"`

	Auth struct {
		JwtToken struct {
			Type           string `mapstructure:"TYPE"`
			Expired        string `mapstructure:"EXPIRED"`
			RefreshExpired string `mapstructure:"REFRESH_EXPIRED"`
		}
	}

	DB struct {
		Mysql struct {
			Host string `mapstructure:"HOST"`
			Port int    `mapstructure:"PORT"`
			Name string `mapstructure:"NAME"`
			User string `mapstructure:"USER"`
			Pass string `mapstructure:"PASS"`
		} `mapstructure:"MYSQL"`
	} `mapstructure:"DB"`

	Cache struct {
		Redis struct {
			Host string `mapstructure:"HOST"`
			Port int    `mapstructure:"PORT"`
			DB   int    `mapstructure:"DB"`
			Pass string `mapstructure:"PASS"`
		}
	}
}

func Get() Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read .env file")
	}

	doOnce.Do(func() {
		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatalln("cannot unmarshaling config")
		}
	})

	return cfg
}
