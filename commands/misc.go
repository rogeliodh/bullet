package commands

import (
	"os"
	"path"

	"github.com/codegangsta/cli"
)

func configFlag() cli.StringFlag {
	configPath := os.Getenv("BULLET_CONFIG")
	if len(configPath) == 0 {
		configPath = path.Join(os.Getenv("HOME"), ".config.bullet")
	}

	flag := cli.StringFlag{
		Name:  "config,c",
		Value: configPath,
		Usage: "The path of your config file",
	}

	return flag
}
