// Package tomlreader implements activities to read and query TOML files
package tomlreader

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

// Update these variables before testing
const (
	filename = "C:\\Users\\lstig\\Downloads\\app\\config.toml"
)

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEvalReadSingleItem(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("filename", filename)
	tc.SetInput("key", "version")

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%s]\n", result)
}

func TestEvalReadNonExistingKey(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("filename", filename)
	tc.SetInput("key", "vers")

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%v]\n", result)
}

func TestEvalReadMultipleItems(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("filename", filename)
	tc.SetInput("key", "items")

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The number of items in the result is:\n[%v]\n", len(result.([]interface{})))
}

func TestEvalReadMultipleItemsWithFilters(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("filename", filename)
	tc.SetInput("key", "items")
	tc.SetInput("filters", "ValueContains(retgits)")

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The number of items in the result is:\n[%v]\n", len(result.([]interface{})))
}

func TestEvalReadMultipleItemsWithMultipleFilters(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("filename", filename)
	tc.SetInput("key", "items")
	tc.SetInput("filters", "ValueContains(retgits)/KeyEquals(type,activity)")

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The number of items in the result is:\n[%v]\n", len(result.([]interface{})))
}
