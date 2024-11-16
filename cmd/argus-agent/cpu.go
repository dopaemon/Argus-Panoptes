package main

import (
	"log"
	"argus/internal/system"
)

func logCPUInfo() {
	cpuInfo, err := system.GetCPUInfo()
	if err != nil {
		log.Fatalf("Error fetching CPU info: %v", err)
	}

	log.Printf("CPU Name: %s", cpuInfo.Name)
	log.Printf("CPU Cores: %d", cpuInfo.Cores)
	log.Printf("CPU Usage: %.2f%%", cpuInfo.Usage)
}

