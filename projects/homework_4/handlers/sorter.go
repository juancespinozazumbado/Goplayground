package handlers

import (
	"sortingPackages/config"
	"sortingPackages/queue"
	"sortingPackages/utils"
	"sync"

	"github.com/sirupsen/logrus"
)

func Sorter(wg *sync.WaitGroup) {
	defer wg.Done()
	utils.Log.Info("Sorter goroutine started")
	for {
		queue.Mutex.Lock()
		if len(queue.UnsortedPackages) > 0 {
			pkg := queue.UnsortedPackages[0]
			queue.UnsortedPackages = queue.UnsortedPackages[1:]
			utils.Log.WithField("ID", pkg.PackageId).Info("Sorting package")

			// Determine the zone
			zone := ""
			for key, value := range config.Zones {
				if value == pkg.Destination {
					zone = key
					break
				}
			}

			if zone == "" {
				utils.Log.WithField("Destination", pkg.Destination).Error("Invalid destination zone")
			} else {
				// Update the status and categorize
				pkg.Status = true
				queue.SortedPackages[zone] = append(queue.SortedPackages[zone], pkg)
				utils.Log.WithFields(logrus.Fields{
					"ID":      pkg.PackageId,
					"Zone":    zone,
					"Status":  "Sorted",
					"Package": pkg,
				}).Info("Package sorted")
			}
		}
		queue.Mutex.Unlock()
	}
}
