package cfg

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
	"tiktokrpc/cmd/api/pkg/constants"
	"tiktokrpc/cmd/api/pkg/errmsg"
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

	constants.QiNiuBucket = Config.GetString("QiNiu.Bucket")
	constants.QiNiuAccessKey = Config.GetString("QiNiu.AccessKey")
	constants.QiNiuSecretKey = Config.GetString("QiNiu.SecretKey")
	constants.QiNiuDomain = Config.GetString("QiNiu.Domain")

	constants.EtcdHost = Config.GetString("Etcd.Host")
	constants.EtcdPort = Config.GetString("Etcd.Port")
	constants.EtcdAddr = fmt.Sprintf("%s:%s", constants.EtcdHost, constants.EtcdPort)

	constants.JaegerHost = Config.GetString("Jaeger.Host")
	constants.JaegerPort = Config.GetString("Jaeger.Port")
	constants.JaegerAddr = fmt.Sprintf("%s:%s", constants.JaegerHost, constants.JaegerPort)

	constants.ServiceHost = Config.GetString("Service.Host")
	constants.ServicePort = Config.GetString("Service.Port")
	constants.WebsocketPort = Config.GetString("Service.WebsocketPort")
	constants.ServiceAddr = fmt.Sprintf("%s:%s", constants.ServiceHost, constants.ServicePort)
	constants.WebsocketAddr = fmt.Sprintf("%s:%s", constants.ServiceHost, constants.WebsocketPort)

	constants.SentinelThreshold = Config.GetFloat64("Sentinel.Threshold")
	constants.SentinelStatIntervalInMs = Config.GetUint32("Sentinel.StatIntervalInMs")

	return nil
}
