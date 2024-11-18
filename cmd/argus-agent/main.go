package main

import (
	"log"
	"argus/internal/system"
	"argus/internal/api"
	"time"
	"github.com/BurntSushi/toml"
	_ "fmt"
)

type Config struct {
	Server struct {
		IP   string `toml:"ip"`
		Port string `toml:"port"`
	} `toml:"server"`
	Security struct {
		APIKey string `toml:"api_key"`
	} `toml:"security"`
}

func main() {
	var config Config
	_, err := toml.DecodeFile("./configs/agent.toml", &config)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	for {
		cpuInfo, err := system.GetCPUInfo()
		if err != nil {
			log.Printf("Error fetching CPU info: %v", err)
		}
		memoryInfo, err := system.GetMemoryInfo()
		if err != nil {
			log.Printf("Error fetching Memory info: %v", err)
		}
		diskInfo, err := system.GetDiskInfo()
		if err != nil {
			log.Printf("Error fetching Disk info: %v", err)
		}

		err = api.SendData(config.Server.IP, config.Server.Port, config.Security.APIKey, cpuInfo, memoryInfo, diskInfo)
		if err != nil {
			log.Printf("Error sending data to API: %v", err)
			time.Sleep(5 * time.Second)
		} else {
			log.Println("Data successfully sent to the server.")
		}

		time.Sleep(10 * time.Second)
	}
}

