// Package addtodate adds a specified number of units to a date.
package addtodate

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	number = "number"
	units  = "units"
	date   = "date"
	result = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-addtodate")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the inputs
	ivNumber := context.GetInput(number).(int)
	ivUnits := context.GetInput(units).(string)
	ivDate := context.GetInput(date).(string)

	date := time.Now()

	if ivDate != "" {
		date, _ = time.Parse("2006-01-02", ivDate)
	}

	switch ivUnits {
	case "days":
		date = date.AddDate(0, 0, 1*ivNumber)
	case "months":
		date = date.AddDate(0, 1*ivNumber, 0)
	case "years":
		date = date.AddDate(1*ivNumber, 0, 0)
	}

	// Set the output value in the context
	context.SetOutput(result, date.Format("2006-01-02"))

	return true, nil
}
