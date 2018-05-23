// Package downloadfile implements a file download for Flogo
package downloadfile

// Imports
import (
	"os"
	"path/filepath"
	"time"

	"github.com/nareix/curl"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants
const (
	ivURL         = "url"
	ivWriteToDisk = "writeToDisk"
	ivFilename    = "filename"
	ivAppend      = "append"
	ivCreate      = "create"
	ovResult      = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-downloadfile")

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
	url := context.GetInput(ivURL).(string)
	writeToDisk := context.GetInput(ivWriteToDisk).(bool)
	filename := context.GetInput(ivFilename).(string)
	append := context.GetInput(ivAppend).(bool)
	create := context.GetInput(ivCreate).(bool)

	// Create new request
	req := curl.Get(url)

	// Set timeouts
	// DialTimeout is the TCP Connection Timeout
	// Timeout is the Download Timeout
	req.DialTimeout(time.Second * 10)
	req.Timeout(time.Second * 30)

	// Specify a progress monitor, otherwise it doesn't work
	req.Progress(func(p curl.ProgressStatus) {}, time.Second)

	// Execute the request and return the result
	res, err := req.Do()
	if err != nil {
		context.SetOutput(ovResult, err.Error())
		return true, err
	}

	if res.StatusCode == 200 {
		if writeToDisk {
			err = writeContentToDisk(filename, res.Body, create, append)
			if err != nil {
				context.SetOutput(ovResult, err.Error())
				return true, err
			}
		} else {
			context.SetOutput(ovResult, res.Body)
			return true, nil
		}
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

func writeContentToDisk(filename string, content string, create bool, append bool) error {
	// Check if the file exists
	_, err := os.Stat(filename)

	if err != nil {
		// If the file doesn't exist but create is set to true, go create the file
		if os.IsNotExist(err) && create {
			createFile(filename)
		} else {
			return err
		}
	}

	var file *os.File

	if append {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	} else {
		file, err = os.OpenFile(filename, os.O_WRONLY, 0600)
	}

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return nil
}
