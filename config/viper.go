package config

import (
	"log"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PORT                string `mapstructure:"PORT"`
	MONGODBURI          string `mapstructure:"MONGODB_URI"`
	MONGODBNAME         string `mapstructure:"MONGODB_NAME"`
	CLOUDINARYAPIKEY    string `mapstructure:"CLOUDINARY_APIKEY"`
	CLOUDINARYSECRET    string `mapstructure:"CLOUDINARY_SECRET"`
	CLOUDINARYCLOUDNAME string `mapstructure:"CLOUDINARY_CLOUDNAME"`
	CLOUDINARYFOLDER    string `mapstructure:"CLOUDINARY_FOLDER"`
	REDISHOST           string `mapstructure:"REDIS_HOST"`
	REDISPORT           string `mapstructure:"REDIS_PORT"`
}

func LoadConfig() (config EnvVars, err error) {
	viper.SetConfigFile(".env")	
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	return
}
