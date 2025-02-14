package boot

import (
	"app-api/config"

	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func SetupConfig() *config.Config {
	// 读取配置信息
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[SetupConfig] viper.ReadInConfig err %s", err)
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("[SetupConfig] viper.Unmarshal err %s", err)
	}

	// 监控配置变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&Config); err != nil {
			log.Fatalf("[SetupConfig] viper.OnConfigChange err %s", err)
		}
	})

	return Config
}
