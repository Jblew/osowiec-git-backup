package util

import "time"

func MeasureDuration(f func() error) (time.Duration, error) {
	start := time.Now()
	err := f()
	ellapsed := time.Since(start)
	return ellapsed, err
}
