// Package trellocard implements activities to create cards in Trello
package trellocard

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

// Update these variables before testing to match your own AWS account
const (
	trelloToken = ""
	trelloKey   = ""
	trelloList  = ""
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

// Test for a card at the top of the list
func TestEvalCreateCardTop(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("token", trelloToken)
	tc.SetInput("appkey", trelloKey)
	tc.SetInput("list", trelloList)
	tc.SetInput("position", "top")
	tc.SetInput("title", "Hello World at the top")
	tc.SetInput("description", "Hello World is actually an awesome description for a card...")

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%s]\n", result)
}

// Test for a card at the bottom of the list
func TestEvalCreateCardBottom(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("token", trelloToken)
	tc.SetInput("appkey", trelloKey)
	tc.SetInput("list", trelloList)
	tc.SetInput("position", "bottom")
	tc.SetInput("title", "Hello World at the bottom")
	tc.SetInput("description", "Hello World is actually an awesome description for a card...")

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%s]\n", result)
}

// Test for a card at the bottom of the list with no description
func TestEvalCreateCardNoDescription(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("token", trelloToken)
	tc.SetInput("appkey", trelloKey)
	tc.SetInput("list", trelloList)
	tc.SetInput("position", "bottom")
	tc.SetInput("title", "Hello World can also have no description")

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%s]\n", result)
}
