// Package randomnumber generates a random number between the minimum and maximum values supplied.
package randomnumber

import (
	"math/rand"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	min    = "min"
	max    = "max"
	result = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-randomnumber")

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
	minvalue := context.GetInput(min).(int)
	maxvalue := context.GetInput(max).(int)

	// Set a random seed
	rand.Seed(time.Now().UTC().UnixNano())

	// Create a random number
	randvalue := minvalue + rand.Intn(maxvalue-minvalue)
	log.Debugf("Created a random number [%s]", randvalue)

	// Set the output value in the context
	context.SetOutput(result, randvalue)

	return true, nil
}
