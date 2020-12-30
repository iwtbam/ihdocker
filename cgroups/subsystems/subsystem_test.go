package subsystems

import (
	"os"
	"path"
	"testing"
)

func TestMemoryCgroup(t *testing.T) {
	subsystem := MemorySubSystem{}
	resConfig := ResourceConfig{
		MemoryLimit: "1000m",
	}

	testCgroup := "testmemlimit"

	if err := subsystem.Set(testCgroup, &resConfig); err != nil {
		t.Fatalf("cgroup fail %v", err)
	}

	stat, _ := os.Stat(path.Join(FindCgroupMountPoint(subsystem.Name()), testCgroup))

	t.Logf("cgroup stats:%v", stat)

	if err := subsystem.Apply(testCgroup, os.Getpid()); err != nil {
		t.Fatalf("cgroup Apply %v", err)
	}

	if err := subsystem.Apply("", os.Getpid()); err != nil {
		t.Fatalf("cgroup Apply %v", err)
	}

	if err := subsystem.Remove(testCgroup); err != nil {
		t.Fatalf("cgroup remove %v", err)
	}
}

func TestCpuCgroup(t *testing.T) {
	subsystem := CpuSubSystem{}
	resConfig := ResourceConfig{
		CpuShare: "512",
	}

	testCgroup := "testmemlimit"

	if err := subsystem.Set(testCgroup, &resConfig); err != nil {
		t.Fatalf("cgroup fail %v", err)
	}

	stat, _ := os.Stat(path.Join(FindCgroupMountPoint(subsystem.Name()), testCgroup))

	t.Logf("cgroup stats:%v", stat)

	if err := subsystem.Apply(testCgroup, os.Getpid()); err != nil {
		t.Fatalf("cgroup Apply %v", err)
	}

	if err := subsystem.Apply("", os.Getpid()); err != nil {
		t.Fatalf("cgroup Apply %v", err)
	}

	if err := subsystem.Remove(testCgroup); err != nil {
		t.Fatalf("cgroup remove %v", err)
	}
}

func TestCpusetCgroup(t *testing.T) {
	subsystem := CpusetSubSystem{}
	resConfig := ResourceConfig{
		CpuSet: "1",
	}

	testCgroup := "testmemlimit"

	if err := subsystem.Set(testCgroup, &resConfig); err != nil {
		t.Fatalf("cgroup fail %v", err)
	}

	stat, _ := os.Stat(path.Join(FindCgroupMountPoint(subsystem.Name()), testCgroup))

	t.Logf("cgroup stats:%v", stat)

	if err := subsystem.Apply(testCgroup, os.Getpid()); err != nil {
		t.Fatalf("cgroup Apply %v", err)
	}

	if err := subsystem.Apply("", os.Getpid()); err != nil {
		t.Fatalf("cgroup Apply %v", err)
	}

	if err := subsystem.Remove(testCgroup); err != nil {
		t.Fatalf("cgroup remove %v", err)
	}
}
