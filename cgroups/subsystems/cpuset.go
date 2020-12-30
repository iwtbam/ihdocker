package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpusetSubSystem struct {
}

func (s *CpusetSubSystem) Set(cgroupName string, res *ResourceConfig) error {

	if res.CpuSet == "" {
		return nil
	}

	if cgroupPath, err := GetCgroupPath(s.Name(), cgroupName, true); err == nil {
		configFile := path.Join(cgroupPath, "cpuset.cpus")
		if err := ioutil.WriteFile(configFile, []byte(res.CpuSet), 0644); err == nil {
			return nil
		} else {
			return fmt.Errorf("set cgroup memory failed : %v", err)
		}
	} else {
		return err
	}
}

func (s *CpusetSubSystem) Apply(cgroupName string, pid int) error {
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

func (s *CpusetSubSystem) Remove(cgroupName string) error {
	if cgroupPath, err := GetCgroupPath(s.Name(), cgroupName, false); err == nil {
		return os.Remove(cgroupPath)
	} else {
		return err
	}
}

func (s *CpusetSubSystem) Name() string {
	return "cpuset"
}
