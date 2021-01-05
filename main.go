package main

import (
	"github.com/urfave/cli"
	"os"

	"cn.iwtbam.ih/settings"

	log "github.com/sirupsen/logrus"
)

func main() {
	app := cli.NewApp()
	app.Name = "ihdocker"
	app.Usage = "create container"

	app.Commands = []cli.Command{
		runCommand,
		initCommand,
		commitCommand,
	}

	app.Before = func(ctx *cli.Context) error {
		log.SetFormatter(&settings.IhSimpleLogFormatter{})
		log.SetOutput(os.Stdout)
		log.SetReportCaller(true)
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
