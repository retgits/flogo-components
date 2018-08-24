// Package readfile implements a file reader for Flogo
package readfile

// Imports
import (
	"io/ioutil"
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants
const (
	ivFilename = "filename"
	ovResult   = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-readfile")

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

	// Check if the file exists
	_, err = os.Stat(filename)
	if err != nil {
		log.Debugf("Error while tryinf to find file: %s", err.Error())
		return false, err
	}

	// Read the file
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Debugf("Error while reading file: %s\n", err.Error())
		return false, err
	}

	context.SetOutput(ovResult, string(fileBytes))
	return true, nil
}
