package amiup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func statusHandlerFunc(w http.ResponseWriter, r *http.Request) {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	status := map[string]interface{}{
		"status":  "OK",
		"uptime":  time.Since(startTime).String(),
		"version": Version,
		"runtime": map[string]interface{}{
			"goroutines": runtime.NumGoroutine(),
			"cpu_count":  runtime.NumCPU(),
			"cgo_calls":  runtime.NumCgoCall(),
			"memory": map[string]interface{}{
				"alloc":       m.Alloc,
				"total_alloc": m.TotalAlloc,
				"sys":         m.Sys,
				"lookups":     m.Lookups,
				"mallocs":     m.Mallocs,
				"frees":       m.Frees,
			},
		},
	}
	d, err := json.Marshal(status)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(d)
	if err != nil {
		fmt.Printf("Error writing response: %s\n", err)
		return
	}
}
