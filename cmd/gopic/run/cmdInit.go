package run

import (
	"io/ioutil"

	"github.com/OSTGO/gopic/pkg/conf"
	"github.com/OSTGO/gopic/pkg/constant"
)

func CmdInit() (string, string) {
	data := conf.Config
	err := ioutil.WriteFile(constant.DEFAULT_CONF_P, []byte(data), 0777)
	if err != nil {
		panic(err)
	}
	return data, constant.DEFAULT_CONF_P
}
