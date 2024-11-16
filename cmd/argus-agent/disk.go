package main

import (
	"log"
	"argus/internal/system"
)

func logDiskInfo() {
	diskInfo, err := system.GetDiskInfo()
	if err != nil {
		log.Fatalf("Error fetching Disk info: %v", err)
	}

	log.Printf("Disk Total: %d bytes", diskInfo.Total)
	log.Printf("Disk Used: %d bytes", diskInfo.Used)
}

