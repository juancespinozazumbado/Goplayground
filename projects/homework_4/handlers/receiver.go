package handlers

import (
	"sortingPackages/models"
	"sortingPackages/queue"
	"sortingPackages/utils"
	"sync"

	"github.com/sirupsen/logrus"
)

// func Receiver(wg *sync.WaitGroup, incoming chan models.Package) {
// 	defer wg.Done()
// 	for pkg := range incoming {
// 		queue.Mutex.Lock()
// 		queue.UnsortedPackages = append(queue.UnsortedPackages, pkg)
// 		queue.Mutex.Unlock()
// 	}
// }

func Receiver(wg *sync.WaitGroup, incoming chan models.Package) {
	defer wg.Done()
	utils.Log.Info("Receiver goroutine started")
	for pkg := range incoming {
		utils.Log.WithFields(logrus.Fields{
			"ID":          pkg.PackageId,
			"SourceName":  pkg.SenderName,
			"Destination": pkg.Destination,
			"Status":      pkg.Status,
		}).Info("Received new package")

		queue.Mutex.Lock()
		queue.UnsortedPackages = append(queue.UnsortedPackages, pkg)
		queue.Mutex.Unlock()
		utils.Log.WithField("UnsortedPackagesCount", len(queue.UnsortedPackages)).Info("Added package to unsorted queue")
	}
}
