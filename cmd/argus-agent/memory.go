package main

import (
	"log"
	"argus/internal/system"
)

func logMemoryInfo() {
	memoryInfo, err := system.GetMemoryInfo()
	if err != nil {
		log.Fatalf("Error fetching Memory info: %v", err)
	}

	log.Printf("Memory Total: %d bytes", memoryInfo.Total)
	log.Printf("Memory Used: %d bytes", memoryInfo.Used)
}

