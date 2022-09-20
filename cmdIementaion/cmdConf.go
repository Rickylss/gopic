package cmdIementaion

import (
	"fmt"

	"github.com/OSTGO/gopic/pkg/conf"
	"github.com/OSTGO/gopic/pkg/constant"
	"github.com/OSTGO/gopic/pkg/utils"
	"golang.org/x/exp/maps"
)

func CmdConf() {
	showConfigAndPath()
	showAllPluginList()
	showActivePluginList()
	showEveryPluginHelp()
}

func showAllPluginList() {
	fmt.Println("Supported Plugins: ", maps.Keys(utils.StroageHelp))
}

func showActivePluginList() {
	fmt.Println("Actived Plugins: ", maps.Keys(utils.StroageMap))
}

func showEveryPluginHelp() {
	for k, v := range utils.StroageHelp {
		fmt.Printf("%s:%s\n", k, v)
	}
}

func showConfigAndPath() {
	fmt.Println("Config Path:" + constant.DEFAULT_CONF_P)
	fmt.Println("Config Content:" + fmt.Sprintln(conf.Viper.AllSettings()))
}
