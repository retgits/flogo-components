package iftttwebhook

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// Update these variables before testing to match your own IFTTT account
const (
	key   = ""
	event = ""
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

func TestEvalOneInput(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("key", key)
	tc.SetInput("event", event)
	tc.SetInput("value1", "simply hello")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("[%s]", result)
}

func TestEvalTwoInputs(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("key", key)
	tc.SetInput("event", event)
	tc.SetInput("value1", "hello")
	tc.SetInput("value2", "world")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("[%s]", result)
}

func TestEvalThreeInputs(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("key", key)
	tc.SetInput("event", event)
	tc.SetInput("value1", "hello")
	tc.SetInput("value2", "beautiful")
	tc.SetInput("value3", "world")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("[%s]", result)
}
