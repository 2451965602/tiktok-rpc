package cfg

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
	"tiktokrpc/cmd/social/pkg/constants"
	"tiktokrpc/cmd/social/pkg/errmsg"
)

var Config *viper.Viper

func Init() error {
	Config = viper.New()
	Config.SetConfigFile("./config/config.yaml")
	Config.SetConfigType("yaml")
	if err := Config.ReadInConfig(); err != nil {
		return errmsg.ConfigMissError
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("tiktok")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := loadConfig()
	if err != nil {

		return err
	}

	go func() {
		Config.WatchConfig()
		Config.OnConfigChange(func(e fsnotify.Event) {
			err := loadConfig()
			if err != nil {
				panic(err)
			}
		})
	}()

	return nil
}

func loadConfig() error {

	constants.MySQLUserName = Config.GetString("MySQL.UserName")
	constants.MySQLPassWord = Config.GetString("MySQL.PassWord")
	constants.MySQLHost = Config.GetString("MySQL.Host")
	constants.MySQLPort = Config.GetString("MySQL.Port")
	constants.MySQLName = Config.GetString("MySQL.Name")
	constants.MySQLDSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", constants.MySQLUserName, constants.MySQLPassWord, constants.MySQLHost, constants.MySQLPort, constants.MySQLName)

	constants.EtcdHost = Config.GetString("Etcd.Host")
	constants.EtcdPort = Config.GetString("Etcd.Port")
	constants.EtcdAddr = fmt.Sprintf("%s:%s", constants.EtcdHost, constants.EtcdPort)

	constants.ServiceHost = Config.GetString("Service.Host")
	constants.ServicePort = Config.GetString("Service.Port")
	constants.ServiceAddr = fmt.Sprintf("%s:%s", constants.ServiceHost, constants.ServicePort)
	return nil
}
