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
	MysqlConfig  *MysqlConfig
}

type OpenaiConfig struct {
	Key string
}

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Db       string
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
	conf.InitMysqlConfig()

	return conf
}

func (c *Config) InitOpenaiConfig() {
	oc := &OpenaiConfig{
		Key: c.viper.GetString("openai.key"),
	}
	c.OpenaiConfig = oc
}

func (c *Config) InitMysqlConfig() {
	mc := &MysqlConfig{
		Username: c.viper.GetString("mysql.username"),
		Password: c.viper.GetString("mysql.password"),
		Host:     c.viper.GetString("mysql.host"),
		Port:     c.viper.GetInt("mysql.port"),
		Db:       c.viper.GetString("mysql.db"),
	}
	c.MysqlConfig = mc
}
