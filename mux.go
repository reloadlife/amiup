package amiup

import (
	"fmt"
	"net/http"
)

func setupMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", statusHandlerFunc)
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			fmt.Printf("Error writing response: %s\n", err)
			return
		}
	})

	return mux
}
