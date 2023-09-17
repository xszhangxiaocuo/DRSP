package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var AppConf = InitConfig()

type Config struct {
	viper        *viper.Viper
	OpenaiConfig *OpenaiConfig
}

type OpenaiConfig struct {
	Key string
}

func InitConfig() *Config {
	v := viper.New()
	conf := &Config{viper: v}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	conf.InitOpenaiConfig()

	return conf
}

func (c *Config) InitOpenaiConfig() {
	oc := &OpenaiConfig{}
	oc.Key = c.viper.GetString("openai.key")
	c.OpenaiConfig = oc
}
