// Package queryparser implements an activity to parse a query string into name/value pairs
package queryparser

import (
	"net/url"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivQuery  = "query"
	ovResult = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-queryparser")

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

	// Get the action
	query := context.GetInput(ivQuery).(string)

	// Parse the string
	m, err := url.ParseQuery(query)
	content := make(map[string]interface{}, 0)
	if err != nil {
		log.Errorf("Error while parsing query string: %s", err.Error())
		return true, err
	}
	for key, val := range m {
		if len(val) == 1 {
			content[key] = val[0]
		} else {
			content[key] = val[0]
		}
	}

	// Set the output value in the context
	context.SetOutput(ovResult, content)

	return true, nil
}
