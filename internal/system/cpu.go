package system

import (
	"runtime"
	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUInfo struct {
	Name       string  `json:"name"`
	Cores      int     `json:"cores"`
	Usage      float64 `json:"usage"`
}

func GetCPUInfo() (CPUInfo, error) {
	usage, err := cpu.Percent(0, false)
	if err != nil {
		return CPUInfo{}, err
	}
	info, err := cpu.Info()
	if err != nil {
		return CPUInfo{}, err
	}
	return CPUInfo{
		Name:  info[0].ModelName,
		Cores: runtime.NumCPU(),
		Usage: usage[0],
	}, nil
}

