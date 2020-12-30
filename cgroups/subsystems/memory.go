package subsystems

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubSystem struct {
}

func (s *MemorySubSystem) Set(cgroupName string, res *ResourceConfig) error {

	if res.MemoryLimit == "" {
		return nil
	}

	log.Infof("set memory")

	if cgroupPath, err := GetCgroupPath(s.Name(), cgroupName, true); err == nil {
		configFile := path.Join(cgroupPath, "memory.limit_in_bytes")
		if err := ioutil.WriteFile(configFile, []byte(res.MemoryLimit), 0644); err == nil {
			return nil
		} else {
			return fmt.Errorf("set cgroup memory failed : %v", err)
		}
	} else {
		return err
	}
}

func (s *MemorySubSystem) Apply(cgroupName string, pid int) error {
	if cgroupPath, err := GetCgroupPath(s.Name(), cgroupName, false); err == nil {
		taskFile := path.Join(cgroupPath, "tasks")
		if err := ioutil.WriteFile(taskFile, []byte(strconv.Itoa(pid)), 0644); err == nil {
			return err
		} else {
			return fmt.Errorf("set cgroup proc failed ：%v", err)
		}
	} else {
		return err
	}
}

func (s *MemorySubSystem) Remove(cgroupName string) error {
	if cgroupPath, err := GetCgroupPath(s.Name(), cgroupName, false); err == nil {
		return os.Remove(cgroupPath)
	} else {
		return err
	}
}

func (s *MemorySubSystem) Name() string {
	return "memory"
}
