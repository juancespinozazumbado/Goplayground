package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sortingPackages/models"
	"sortingPackages/queue"
	"sortingPackages/utils"

	"github.com/sirupsen/logrus"
)

var IncomingChannel = make(chan models.Package, 10)

func AddPackage(w http.ResponseWriter, r *http.Request) {
	var pkg models.Package
	if err := json.NewDecoder(r.Body).Decode(&pkg); err != nil {
		utils.Log.WithError(err).Error("Failed to decode package payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	pkg.PackageId = utils.GenerateID()
	IncomingChannel <- pkg

	utils.Log.WithFields(logrus.Fields{
		"ID":          pkg.PackageId,
		"SourceName":  pkg.SenderName,
		"Destination": pkg.Destination,
		"Status":      pkg.Status,
	}).Info("Package added to incoming channel")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Package received")
}

func GetPackages(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("zone")
	queue.Mutex.Lock()
	defer queue.Mutex.Unlock()

	if query != "" {
		packages, exists := queue.SortedPackages[query]
		if !exists {
			utils.Log.WithField("Zone", query).Error("Zone not found")
			http.Error(w, "Zone not found", http.StatusNotFound)
			return
		}
		utils.Log.WithField("Zone", query).Info("Fetched sorted packages for zone")
		json.NewEncoder(w).Encode(packages)
		return
	}

	utils.Log.Info("Fetched all unsorted packages")
	json.NewEncoder(w).Encode(queue.UnsortedPackages)
}
