package api

import (
	"encoding/json"
	"net/http"

	"github.com/FOXHOUND0x/ragnarok/internal/monitor"
)

func HandleContainers(w http.ResponseWriter, r *http.Request) {
	containers, err := monitor.ListContainers()
	if err != nil {
		http.Error(w, "Error Listing Containers", http.StatusInternalServerError)
		return
	}

	resp := make([]map[string]string, len(containers))
	for x, container := range containers {
		healthStats, err := monitor.GetContainerHealth(container.ID)
		if err != nil {
			healthStats = "unknown"
		}
		resp[x] = map[string]string{
			"id":     container.ID,
			"name":   container.Names[0],
			"status": container.Status,
			"health": healthStats,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
