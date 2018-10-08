// Package stoplight implements an activity that looks at the incoming data elements
package stoplight

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivWhitelistArray = "whitelistArray"
	ivWhitelist      = "whitelist"
	ivKey            = "key"
	ivValue          = "value"
	ovResult         = "result"
	red              = "RED"
	green            = "GREEN"
)

// log is the default package logger
var log = logger.GetLogger("activity-stoplight")

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
	// Get data elements from the context
	whitelistArray := context.GetInput(ivWhitelistArray).(map[string]interface{})
	whitelist := context.GetInput(ivWhitelist).(string)
	key := context.GetInput(ivKey).(string)
	value := context.GetInput(ivValue).(string)

	// Check if the whitelist name exists in the array
	if val, ok := whitelistArray[whitelist]; ok {
		// Check if the key exists
		testArray := val.([]interface{})

		for _, row := range testArray {
			testMap := row.(map[string]interface{})
			if testMap["name"].(string) == value && testMap["id"].(string) == key {
				context.SetOutput(ovResult, green)
				return true, nil
			}
		}

		context.SetOutput(ovResult, red)
	} else {
		return false, fmt.Errorf("whitelist [%s] does not exist in array", whitelist)
	}

	return true, nil
}
