// Package envkey implements an activity to get environment variables or the provided fallback value
package envkey

import (
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	ivEnvKey   = "envkey"
	ivFallback = "fallback"

	ovResult = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-envkey")

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
	key := context.GetInput(ivEnvKey).(string)
	fallback := context.GetInput(ivFallback).(string)

	// Set the output
	context.SetOutput(ovResult, GetEnvKey(key, fallback))

	return true, nil
}

// GetEnvKey tries to get the specified key from the OS environment and returns either the
// value or the fallback that was provided
func GetEnvKey(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
