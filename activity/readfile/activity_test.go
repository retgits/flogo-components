// Package readfile implements a file reader for Flogo
package readfile

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

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

func TestEvalReadNonExistingFile(t *testing.T) {

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

	// Execute the activity
	_, err := act.Eval(tc)

	assert.Error(t, err, "")
}

func TestEvalReadExistingFile(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("filename", "./activity.json")

	// Execute the activity
	_, err := act.Eval(tc)

	assert.NoError(t, err, "")

	//check result attr
	result := tc.GetOutput("result")
	assert.Contains(t, result, "github.com/retgits/flogo-components/activity/readfile")
}
