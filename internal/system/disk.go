package system

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type DiskInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

func GetDiskInfo() (DiskInfo, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		return DiskInfo{}, err
	}
	return DiskInfo{
		Total: usage.Total,
		Used:  usage.Used,
	}, nil
}

