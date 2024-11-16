package api

type Payload struct {
	CPU    interface{} `json:"cpu"`
	Memory interface{} `json:"memory"`
	Disk   interface{} `json:"disk"`
}

func NewPayload(cpu, memory, disk interface{}) Payload {
	return Payload{
		CPU:    cpu,
		Memory: memory,
		Disk:   disk,
	}
}

