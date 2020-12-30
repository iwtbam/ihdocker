package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

const MOUNT_INFO_PATH = "/proc/self/mountinfo"

func FindCgroupMountPoint(subsystem string) string {
	f, err := os.Open(MOUNT_INFO_PATH)
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(text, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	return ""
}

func GetCgroupPath(subsystem string, cgroupName string, autoCreate bool) (string, error) {

	cgroupRoot := FindCgroupMountPoint(subsystem)
	cgroupPath := path.Join(cgroupRoot, cgroupName)

	if _, err := os.Stat(cgroupPath); err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(cgroupPath, 0755); err == nil {
				return cgroupPath, nil
			} else {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}

		return cgroupPath, nil
	}

	return "", nil
}
