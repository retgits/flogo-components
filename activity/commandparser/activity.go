// Package commandparser implements activities to read and parse commandline arguments
package commandparser

// Imports
import (
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants
const (
	ivCommandString = "commandString"
	ovResult        = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-commandparser")

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
	commandString := context.GetInput(ivCommandString).(string)

	// Parse the command line into separate arguments
	array, err := parseCommandLine(commandString)
	if err != nil {
		log.Error(err.Error())
		return true, err
	}

	// Prepare a datamap for the response
	datamap := make(map[string]string)

	// Loop over the separate arguments to create key/value pairs
	for i := 0; i < len(array); i++ {
		datamap[array[i]] = array[i+1]
		i++
	}

	// Set the output value in the context
	context.SetOutput(ovResult, datamap)

	return true, nil
}

func parseCommandLine(command string) ([]string, error) {
	var args []string
	state := "start"
	current := ""
	quote := "\""
	escapeNext := true

	for i := 0; i < len(command); i++ {
		c := command[i]

		if state == "quotes" {
			if string(c) != quote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = "start"
			}
			continue
		}

		if escapeNext {
			current += string(c)
			escapeNext = false
			continue
		}

		if c == '\\' {
			escapeNext = true
			continue
		}

		if c == '"' || c == '\'' {
			state = "quotes"
			quote = string(c)
			continue
		}

		if state == "arg" {
			if c == ' ' || c == '\t' {
				// Strip off the leading --
				args = append(args, strings.Replace(current, "--", "", 1))
				current = ""
				state = "start"
			} else {
				current += string(c)
			}
			continue
		}

		if c != ' ' && c != '\t' {
			state = "arg"
			current += string(c)
		}
	}

	if state == "quotes" {
		return []string{}, fmt.Errorf("Unclosed quote in command line: %s", command)
	}

	if current != "" {
		// Strip off the leading -
		args = append(args, strings.Replace(current, "-", "", 1))
		// Assume it is a boolean flag set to true
		args = append(args, "true")
	}

	return args, nil
}
