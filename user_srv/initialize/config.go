package initialize

import (
	"fmt"
	"user_srv/global"

	"github.com/spf13/viper"
)

func Config() {
	runmode := "dev"
	configFileName := fmt.Sprintf("config_%s.yaml", runmode)

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.Conf); err != nil {
		panic(err)
	}
}
