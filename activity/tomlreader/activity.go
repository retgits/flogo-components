// Package tomlreader implements activities to read and query TOML files
package tomlreader

// Imports
import (
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	toml "github.com/pelletier/go-toml"
)

// Constants
const (
	ivFilename = "filename"
	ivKey      = "key"
	ovResult   = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-tomlreader")

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
	filename := context.GetInput(ivFilename).(string)
	key := context.GetInput(ivKey).(string)

	// Check if the file exists
	_, err = os.Stat(filename)

	if err != nil {
		log.Error(err.Error())
		return true, err
	}

	// Read the file
	config, err := toml.LoadFile(filename)

	if err != nil {
		log.Error(err.Error())
		return true, err
	}

	// Find the keys
	queryResult := config.Get(key)
	if queryResult == nil {
		// Set the output value in the context
		context.SetOutput(ovResult, nil)
		return true, nil
	}
	resultArray := queryResult.([]*toml.Tree)

	// Prepare a result structure
	datamap := make([]interface{}, len(resultArray))

	// Loop over the queryResult and make a proper interface from it
	for idx, val := range resultArray {
		datamap[idx] = val.ToMap()
	}

	// Set the output value in the context
	context.SetOutput(ovResult, datamap)

	return true, nil
}
