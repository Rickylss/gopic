package plugin

import (
	"os"
	"path"

	"github.com/OSTGO/gopic/pkg/conf"
	"github.com/OSTGO/gopic/pkg/utils"
	"github.com/lithammer/dedent"
)

type LocalStorage struct {
	*utils.MetaStorage
}

const (
	pluginName = "local"
)

var localConfig map[string]interface{}

func (l *LocalStorage) Upload(im *utils.Image) (string, error) {
	localPath := localConfig["path"].(string)
	return localPath, uploadPictureToLocal(localPath, im)
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{utils.NewMetaStorage()}
}

func uploadPictureToLocal(localPath string, im *utils.Image) error {
	if err := os.MkdirAll(path.Join(localPath, im.FolderName), 0755); err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(localPath, im.OutSuffix), []byte(im.OutBase64), 0755); err != nil {
		return err
	}

	return nil
}

func init() {
	utils.StroageHelp[pluginName] = localHelp()
	localConfig = conf.Viper.GetStringMap(pluginName)
	if localConfig == nil {
		return
	}
	active := localConfig["active"]
	if active == nil {
		return
	}
	if active == true {
		utils.StroageMap[pluginName] = NewLocalStorage()
	}
}

func localHelp() string {
	return dedent.Dedent(`
    local plugin need this parameters:
    active: false or true
    path:   like /home/ubuntu/gopic-test/
	`)
}
