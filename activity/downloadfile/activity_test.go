// Package downloadfile implements a file download for Flogo
package downloadfile

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

const url string = ""

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

func TestEvalDownloadFile(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("url", url)
	tc.SetInput("writeToDisk", false)
	tc.SetInput("filename", "./tmp/data.txt")
	tc.SetInput("append", false)
	tc.SetInput("create", true)

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%s]\n", result)
}

func TestEvalDownloadFileToDisk(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("url", url)
	tc.SetInput("writeToDisk", true)
	tc.SetInput("filename", "./tmp/data.txt")
	tc.SetInput("append", false)
	tc.SetInput("create", true)

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%s]\n", result)
}
