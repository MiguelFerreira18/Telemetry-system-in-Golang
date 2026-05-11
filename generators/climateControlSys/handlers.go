package main

import (
	"net/http"
	"os"
)

func router() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if !isHealthy {
				global.setAlgo(Healthy{}, 0)
				isHealthy = true
				global.Logger.Info("Climate Control System generator switched to healthy mode")
			} else {
				global.Logger.Info("System already in healthy mode")
			}
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	router.HandleFunc("/unhealthy", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if isHealthy {
				global.setAlgo(Unhealthy{}, 0)
				isHealthy = false
				global.Logger.Info("Climate Control System generator switched to unhealthy mode")
			}
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	router.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			global.Logger.Info("System Shutting Down")
			os.Exit(ExitOk)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	return &http.Server{
		Addr:    ":9093",
		Handler: router,
	}
}
