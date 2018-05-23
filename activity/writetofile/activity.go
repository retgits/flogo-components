// Package writetofile implements a file writer for Flogo
package writetofile

// Imports
import (
	"os"
	"path/filepath"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants
const (
	ivFilename = "filename"
	ivContent  = "content"
	ivAppend   = "append"
	ivCreate   = "create"
	ovResult   = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-writetofile")

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
	content := context.GetInput(ivContent).(string)
	append := context.GetInput(ivAppend).(bool)
	create := context.GetInput(ivCreate).(bool)

	// Check if the file exists
	_, err = os.Stat(filename)

	if err != nil {
		// If the file doesn't exist but create is set to true, go create the file
		if os.IsNotExist(err) && create {
			createFile(filename)
		} else {
			context.SetOutput(ovResult, err.Error())
			return true, err
		}
	}

	var file *os.File

	if append {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	} else {
		file, err = os.OpenFile(filename, os.O_WRONLY, 0600)
	}

	if err != nil {
		context.SetOutput(ovResult, err.Error())
		return true, err
	}

	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		context.SetOutput(ovResult, err.Error())
		return true, err
	}

	context.SetOutput(ovResult, "OK")
	return true, nil
}

func createFile(filename string) error {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
