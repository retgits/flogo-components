// Package tomlreader implements activities to read and query TOML files
package tomlreader

// Imports
import (
	"os"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	toml "github.com/pelletier/go-toml"
)

// Constants
const (
	ivFilename = "filename"
	ivKey      = "key"
	ivFilters  = "filters"
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
	filters := context.GetInput(ivFilters).(string)

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

	// Filter the results
	if len(filters) > 0 {
		filterArray := strings.Split(filters, "/")
		for _, filter := range filterArray {
			if strings.HasPrefix(filter, "ValueContains") {
				value := strings.Replace(strings.Split(filter, "(")[1], ")", "", -1)
				datamap = mapValueContains(datamap, value)
			} else if strings.HasPrefix(filter, "KeyEquals") {
				params := strings.Replace(strings.Split(filter, "(")[1], ")", "", -1)
				vals := strings.Split(params, ",")
				datamap = mapKeyEquals(datamap, vals[0], vals[1])
			}
		}
	}

	// Set the output value in the context
	context.SetOutput(ovResult, datamap)

	return true, nil
}

func mapKeyEquals(datamap []interface{}, key string, value interface{}) []interface{} {
	tempmap := make([]interface{}, 0)

	for _, val := range datamap {
		item := val.(map[string]interface{})
		if _, ok := item[key]; ok {
			if item[key] == value {
				tempmap = append(tempmap, item)
			}
		}
	}

	return tempmap
}

func mapValueContains(datamap []interface{}, value string) []interface{} {
	tempmap := make([]interface{}, 0)

	for _, val := range datamap {
		item := val.(map[string]interface{})
		for k := range item {
			if strings.Contains(item[k].(string), value) {
				tempmap = append(tempmap, item)
				break
			}
		}
	}

	return tempmap
}
