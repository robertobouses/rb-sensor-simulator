package domain

import "time"

type SensorConfig struct {
	SamplingInterval time.Duration
	AlertThreshold   float64
	Unit             string
	Enabled          bool
}
