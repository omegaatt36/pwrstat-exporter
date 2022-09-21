package gopwrstat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstruct(t *testing.T) {
	s := assert.New(t)

	sample := `
The UPS information shows as following:

	Properties:
			Model Name................... CP1000PFCLCDa
			Firmware Number.............. CR01803BBI11
			Rating Voltage............... 120 V
			Rating Power................. 600 Watt(1000 VA)

	Current UPS status:
			State........................ Normal
			Power Supply by.............. Utility Power
			Utility Voltage.............. 114 V
			Output Voltage............... 114 V
			Battery Capacity............. 100 %
			Remaining Runtime............ 35 min.
			Load......................... 114 Watt(19 %)
			Line Interaction............. None
			Test Result.................. Passed at 2022/09/21 20:44:29
			Last Power Event............. None
			`

	actual := construct(sample)
	s.Len(actual.Status, 14)
	actualBatteryCapacity, ok := actual.Status["Battery Capacity"]
	s.True(ok)
	s.Equal("100 %", actualBatteryCapacity)
}
