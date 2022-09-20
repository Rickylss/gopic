package conf

import (
	"fmt"

	"github.com/OSTGO/gopic/pkg/constant"
	"github.com/spf13/viper"
)

var Viper viper.Viper

func init() {
	Viper = *viper.New()
	Viper.SetConfigFile(constant.DEFAULT_CONF_P)
	if err := Viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
