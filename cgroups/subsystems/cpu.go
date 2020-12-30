package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpuSubSystem struct {
}

func (s *CpuSubSystem) Set(cgroupName string, res *ResourceConfig) error {

	if res.CpuShare == "" {
		return nil
	}

	if cgroupPath, err := GetCgroupPath(s.Name(), cgroupName, true); err == nil {
		configFile := path.Join(cgroupPath, "cpu.shares")
		if err := ioutil.WriteFile(configFile, []byte(res.CpuShare), 0644); err == nil {
			return nil
		} else {
			return fmt.Errorf("set cgroup cpu shared failed : %v", err)
		}
	} else {
		return err
	}
}

func (s *CpuSubSystem) Apply(cgroupName string, pid int) error {
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

func (s *CpuSubSystem) Remove(cgroupName string) error {
	if cgroupPath, err := GetCgroupPath(s.Name(), cgroupName, false); err == nil {
		return os.Remove(cgroupPath)
	} else {
		return err
	}
}

func (s *CpuSubSystem) Name() string {
	return "cpu"
}
