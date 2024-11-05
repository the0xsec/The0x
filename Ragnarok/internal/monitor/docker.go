package monitor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

// Container represents a simplified Docker container structure
type Container struct {
	ID     string   `json:"Id"`
	Names  []string `json:"Names"`
	Status string   `json:"Status"`
	Health string   `json:"Health"`
}

// ListContainers lists all Docker containers using the Docker CLI
func ListContainers() ([]Container, error) {
	cmd := exec.Command("docker", "ps", "--format", "{{json .}}")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	var containers []Container
	for _, line := range bytes.Split(out.Bytes(), []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		var container Container
		if err := json.Unmarshal(line, &container); err != nil {
			return nil, fmt.Errorf("failed to parse container info: %w", err)
		}
		containers = append(containers, container)
	}

	return containers, nil
}

// GetContainerHealth retrieves the health status of a given container using the Docker CLI
func GetContainerHealth(containerID string) (string, error) {
	cmd := exec.Command("docker", "inspect", "--format", "{{.State.Health.Status}}", containerID)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to inspect container: %w", err)
	}

	return out.String(), nil
}
