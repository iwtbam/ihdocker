package main

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
	"path/filepath"
)

func commitContainer(imageName string) {
	mntURL := "/root/mnt"
	imageTar := filepath.Join("/root/", imageName+".tar")
	log.Infof("tar : %s", imageTar)
	if _, err := exec.Command("tar", "-czf", imageTar, "-C", mntURL, ".").CombinedOutput(); err != nil {
		log.Errorf("Tar folder %s error : %v", mntURL, err)
	}
}
