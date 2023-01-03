package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	ServerPort  string `mapstructure:"ServerPort"`
	Environment string `mapstructure:"Environment"`
	LogOutput   string `mapstructure:"LogOutput"`
	DBUsername  string `mapstructure:"DBUsername"`
	DBPassword  string `mapstructure:"DBPassword"`
	DBHost      string `mapstructure:"DBHost"`
	DBPort      string `mapstructure:"DBPort"`
	DBName      string `mapstructure:"DBName"`
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	// viper.AddConfigPath("../")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("☠️ Env config file not found: ", err)
		} else {
			log.Fatal("☠️ Env config file error: ", err)
		}
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	log.Printf("%+v \n", env)
	return env
}
