package scheduller

import "time"

func Schedule(f func(), interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		<-ticker.C
		f()
	}
}
