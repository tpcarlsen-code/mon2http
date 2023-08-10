package monitors

import "mon2http/src/sources"

const uptimeType = "uptime"

func init() {
	registerMonitor(NewUptime(sources.NewUptimeSource()))
}

type Uptime struct {
	BaseMonitor
}

func NewUptime(s *sources.Uptime) *Uptime {
	m := Uptime{}
	m.source = s
	m.typeString = memoryUsageType
	return &m
}
