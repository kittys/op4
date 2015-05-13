// This package is for Hana operation 4 software

package main

import (
	"fmt"
	"time"
)

type status uint8

// Equipment can be machine, jig fixture or spare part
type Equipment struct {
	Name       string
	Status     string
	Duration   time.Duration
	RegistDate time.Time

	numDown       int
	totalUptime   time.Duration
	totalDowntime time.Duration
}

func (e *Equipment) Stringer() string {
	return fmt.Sprintf("%s is %s for %s", e.Name, e.Status, e.Duration)
}

func (e *Equipment) Age() time.Duration {
	return time.Since(e.RegistDate)
}

// Mean time between failure = total_uptime / num_down
func (e *Equipment) MTBF() time.Duration {
	if e.numDown > 0 {
		return e.totalUptime / time.Duration(e.numDown)
	} else {
		return e.totalUptime
	}
}

// Mean time to repair = total_downtime / num_down
func (e *Equipment) MTTR() time.Duration {
	if e.numDown > 0 {
		return e.totalDowntime / time.Duration(e.numDown)
	} else {
		return e.totalDowntime
	}
}

// Overall Equipment Effectiveness refer to
// http://www.tpmconsulting.org/menu3_show.php?id=12 and http://www.oee.com/calculating-oee.html
// This simple version will calculate from availability only
func (e *Equipment) OEE() float64 {
	return e.Availability()
}

// Availability is percentage of time machine running vs time planned running
func (e *Equipment) Availability() float64 {
	var operatingTime float64 = 42
	var loadingTime float64 = 42
	return operatingTime / loadingTime
}
