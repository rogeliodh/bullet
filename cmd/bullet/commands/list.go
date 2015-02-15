package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/bullet/cmd/bullet/app"
	"github.com/mitsuse/bullet/pushbullet"
)

func NewListCommand() cli.Command {
	command := cli.Command{
		Name:      "list",
		ShortName: "l",
		Usage:     "List devices that can be pushed to",
		Action:    actionList,

		Flags: []cli.Flag{
			configFlag(),
		},
	}

	return command
}

func actionList(ctx *cli.Context) {
	configPath := ctx.String("config")

	config, err := app.LoadConfigPath(configPath)
	if err != nil {
		app.PrintError(err)
		return
	}

	pb := pushbullet.New(config.Token())

	res, err := pb.GetDevices()
	if err != nil {
		app.PrintError(err)
		return
	}

	for _, device := range res.Devices {
		if !device.Pushable {
			continue
		}

		fmt.Println(device.Nickname)
	}
}
