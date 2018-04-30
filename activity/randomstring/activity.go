// Package randomstring generates a random string consisting with the length you specify.
package randomstring

import (
	"math/rand"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	length = "length"
	result = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-randomstring")

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
	stringLength := context.GetInput(length).(int)

	// Set a random seed
	rand.Seed(time.Now().UTC().UnixNano())

	// Create a random string
	bytes := make([]byte, stringLength)
	for i := 0; i < stringLength; i++ {
		bytes[i] = byte(randInt(65, 90))
	}

	log.Debugf("Created a random string [%s]", string(bytes))

	// Set the output value in the context
	context.SetOutput(result, string(bytes))

	return true, nil
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
