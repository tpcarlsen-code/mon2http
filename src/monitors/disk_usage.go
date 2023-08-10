package monitors

import "mon2http/src/sources"

const diskUsageType = "disk_usage"
const diskUsageLimit = 80 // milli-percent

func init() {
	registerMonitor(NewDiskUsage(sources.NewDiskUsage()))
}

type DiskUsage struct {
	BaseMonitor
}

func NewDiskUsage(s *sources.DiskUsage) *DiskUsage {
	m := DiskUsage{}
	m.source = s
	m.limit = diskUsageLimit
	m.typeString = diskUsageType
	return &m
}
