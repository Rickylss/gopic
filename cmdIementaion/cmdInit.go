package cmdIementaion

import (
	"io/ioutil"

	"github.com/OSTGO/gopic/pkg/constant"
	"github.com/OSTGO/gopic/staic"
)

func CmdInit() (string, string) {
	data := staic.Config
	configPath := constant.DEFAULT_CONF_P
	err := ioutil.WriteFile(configPath, []byte(data), 0777)
	if err != nil {
		panic(err)
	}
	return data, configPath
}
