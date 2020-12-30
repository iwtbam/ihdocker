package cgroups

import (
	"cn.iwtbam.ih/cgroups/subsystems"
	log "github.com/sirupsen/logrus"
)

type CgroupManager struct {
	Path     string
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{Path: path}
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subsystem := range subsystems.SubsystemsIns {
		if err := subsystem.Apply(c.Path, pid); err != nil {
			return err
		}
	}
	return nil
}

func (c *CgroupManager) Set(res *subsystems.ResourceConfig) {
	for _, subsystem := range subsystems.SubsystemsIns {
		if err := subsystem.Set(c.Path, res); err != nil {
			log.Warnf("set cgroup failed : %v", err)
		}
	}
}

func (c *CgroupManager) Destroy() error {
	for _, subsystem := range subsystems.SubsystemsIns {
		if err := subsystem.Remove(c.Path); err != nil {
			log.Warnf("remove %s cgroup failed : %v", subsystem.Name(), err)
		}
	}
	return nil
}
