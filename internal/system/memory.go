package system

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

func GetMemoryInfo() (MemoryInfo, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return MemoryInfo{}, err
	}
	return MemoryInfo{
		Total: vm.Total,
		Used:  vm.Used,
	}, nil
}

