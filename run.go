package main

import (
	"cn.iwtbam.ih/cgroups"
	"cn.iwtbam.ih/cgroups/subsystems"
	"cn.iwtbam.ih/container"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Run(tty bool, cmdArray []string, res *subsystems.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)

	if parent == nil {
		log.Errorf("New parent process error")
		return
	}

	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	log.Infof("parent start!")

	cgroupManager := cgroups.NewCgroupManager("ihdocker-cgroups")
	defer cgroupManager.Destroy()

	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)
	sendInitCommand(cmdArray, writePipe)
	parent.Wait()
	log.Infof("parent exit")

}

func sendInitCommand(cmdArray []string, writePipe *os.File) {
	command := strings.Join(cmdArray, " ")
	log.Infof("command all is : %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
