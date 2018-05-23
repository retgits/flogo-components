// Package writetofile implements a file writer for Flogo
package writetofile

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

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

func TestEvalCreateNewFile(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("filename", "./tmp/data.txt")
	tc.SetInput("content", "Hello World!!")
	tc.SetInput("append", false)
	tc.SetInput("create", true)

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%s]\n", result)
}

func TestEvalAppendToFile(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("filename", "./tmp/data.txt")
	tc.SetInput("content", "\nA second Hello World!!")
	tc.SetInput("append", true)
	tc.SetInput("create", false)

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%s]\n", result)
}

func TestEvalAppendToNonExisting(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("filename", "./tmp/bla.txt")
	tc.SetInput("content", "Hello World!!")
	tc.SetInput("append", false)
	tc.SetInput("create", false)

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%s]\n", result)
}
