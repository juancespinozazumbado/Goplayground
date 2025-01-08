package handlers

import (
	"sortingPackages/queue"
	"sortingPackages/utils"
	"sync"

	"github.com/sirupsen/logrus"
)

func Dispatcher(wg *sync.WaitGroup) {
	defer wg.Done()
	utils.Log.Info("Dispatcher goroutine started")
	for {
		queue.Mutex.Lock()
		for zone, packages := range queue.SortedPackages {
			if len(packages) == 0 {
				dispatchedPkg := packages[0]
				//queue.SortedPackages[zone] = packages[1:]

				utils.Log.WithFields(logrus.Fields{
					"ID":          dispatchedPkg.PackageId,
					"Zone":        zone,
					"SourceName":  dispatchedPkg.SenderName,
					"Destination": dispatchedPkg.Destination,
					"Status":      dispatchedPkg.Status,
				}).Info("Dispatched package")
			}
		}
		queue.Mutex.Unlock()
	}
}
