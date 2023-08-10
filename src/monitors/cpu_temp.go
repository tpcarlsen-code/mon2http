package monitors

import (
	"mon2http/src/sources"
)

const cpuTempType = "cpu_temp"
const cpuTempLimit = 75 // degrees C

func init() {
	registerMonitor(NewCPUTemp(sources.NewCPUTempSource()))
}

type CPUTemp struct {
	BaseMonitor
}

func NewCPUTemp(source *sources.CPUTemp) *CPUTemp {
	m := CPUTemp{}
	m.limit = cpuTempLimit
	m.source = source
	m.typeString = cpuTempType
	return &m
}
