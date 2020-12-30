package main

import (
	"cn.iwtbam.ih/cgroups/subsystems"
	"cn.iwtbam.ih/container"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "run container",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}

		var cmdArray []string

		for _, arg := range ctx.Args() {
			cmdArray = append(cmdArray, arg)
		}

		tty := ctx.Bool("it")

		resConf := &subsystems.ResourceConfig{
			MemoryLimit: ctx.String("m"),
			CpuSet:      ctx.String("cpuset"),
			CpuShare:    ctx.String("cpushare"),
		}

		Run(tty, cmdArray, resConf)
		return nil
	},

	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},

		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},

		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},

		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "init container",
	Action: func(ctx *cli.Context) error {
		log.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}
