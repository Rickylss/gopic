package cmdIementaion

import (
	"fmt"

	"github.com/OSTGO/gopic/pkg/conf"
	"github.com/OSTGO/gopic/pkg/constant"
	"github.com/OSTGO/gopic/pkg/utils"
	"golang.org/x/exp/maps"
)

func CmdConf() string {
	return showConfigAndPath() + "\n" + showAllPluginList() + "\n" + showActivePluginList() + "\n" + showEveryPluginHelp()
}

func showAllPluginList() string {
	return fmt.Sprintln("Supported Plugins:", maps.Keys(utils.StroageHelp))
}

func showActivePluginList() string {
	return fmt.Sprintln("Actived Plugins: ", maps.Keys(utils.StroageMap))
}

func showEveryPluginHelp() string {
	return fmt.Sprintln(utils.StroageHelp)
}

func showConfigAndPath() string {
	return "Config Path:" + constant.DEFAULT_CONF_P + "Config Content:" + fmt.Sprintln(conf.Viper.AllSettings())
}
