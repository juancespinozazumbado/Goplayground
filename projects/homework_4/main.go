package main

import (
	"net/http"
	"sortingPackages/api"
	"sortingPackages/handlers"
	"sortingPackages/utils"
	"sync"
)

func main() {

	utils.Log.Info("Starting Package Sorter application")
	var wg sync.WaitGroup

	// Start goroutines
	wg.Add(3)
	go handlers.Receiver(&wg, api.IncomingChannel)
	go handlers.Sorter(&wg)
	go handlers.Dispatcher(&wg)

	// Start HTTP server
	http.HandleFunc("/add-package", api.AddPackage)
	http.HandleFunc("/get-packages", api.GetPackages)

	utils.Log.Info("Starting HTTP server on port 8080")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			utils.Log.WithError(err).Fatal("HTTP server encountered an error")
		}
	}()

	wg.Wait()
	utils.Log.Info("Shutting down Package Sorter application")
}
