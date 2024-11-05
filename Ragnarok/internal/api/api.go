package api

import (
	"net/http"
)

func RoutesConfig() {
	http.HandleFunc("/containers", HandleContainers)
}
