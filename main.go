package amiup

import (
	"fmt"
	"github.com/sarulabs/di"
	"net/http"
	"time"
)

var (
	startTime = time.Now()
	Version   = "v1.0.0" // this will change using ldflags
)

func ServiceStart(Addr string) {
	mux := setupMux()
	err := http.ListenAndServe(Addr, mux)
	if err != nil {
		fmt.Printf("Error starting status service: %s\n", err)
		return
	}
}

var Service = "status"

func DiService(Addr string) *di.Def {
	return &di.Def{
		Name: Service,
		Build: func(ctn di.Container) (interface{}, error) {
			mux := setupMux()
			go func() {
				err := http.ListenAndServe(Addr, mux)
				if err != nil {
					fmt.Printf("Error starting status service: %s\n", err)
				}
			}()
			return mux, nil
		},
	}
}
