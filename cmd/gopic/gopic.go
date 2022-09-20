package gopic

import (
	"os"

	"github.com/OSTGO/gopic/cmd/gopic/cmd"
	log "github.com/sirupsen/logrus"
)

func Run() error {
	return cmd.NewCommand().Execute()
}

func Main() {
	// let's explicitly set stdout
	log.SetOutput(os.Stdout)
	// this formatter is the default, but the timestamps output aren't
	// particularly useful, they're relative to the command start
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
	if err := Run(); err != nil {
		os.Exit(1)
	}
}
