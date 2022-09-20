package run

import (
	"fmt"
	"testing"
)

func TestCmdUpload(t *testing.T) {
	uf := &UploadFlags{
		Path:        "../1.png",
		StorageList: []string{"qiniu", "github"},
		AllStorage:  false,
		NameReserve: false,
		OutFormat:   "qiniu",
	}
	args := []string{"../2.png"}
	uo := uf.NewUploadOptions(nil, args)

	outs := uo.CmdUpload()
	fmt.Println(outs)
	fmt.Println("end")
}
