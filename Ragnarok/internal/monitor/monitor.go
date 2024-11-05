package monitor

import (
	"fmt"
)

// Monitor is responsible for checking the health of Docker containers
type Monitor struct{}

// NewMonitor creates a new Monitor instance
func NewMonitor() *Monitor {
	return &Monitor{}
}

// DisplayContainerHealth displays the health status of all containers
func (m *Monitor) DisplayContainerHealth() {
	containers, err := ListContainers()
	if err != nil {
		fmt.Println("Error listing containers:", err)
		return
	}

	for _, container := range containers {
		healthStatus, err := GetContainerHealth(container.ID)
		if err != nil {
			fmt.Printf("Error checking health for container %s: %v\n", container.ID, err)
			continue
		}
		fmt.Printf("Container %s (%s) is %s\n", container.Names[0], container.ID[:12], healthStatus)
	}
}
