// Package amazonsqssend sends a message using Amazon SQS
package amazonsqssend

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
	awsAccessKeyID     = ""
	awsSecretAccessKey = ""
	awsDefaultRegion   = ""
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

// Test for a single message
func TestEvalSendMessage(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("AWSAccessKeyID", awsAccessKeyID)
	tc.SetInput("AWSSecretAccessKey", awsSecretAccessKey)
	tc.SetInput("AWSDefaultRegion", awsDefaultRegion)
	tc.SetInput("QueueUrl", "")
	tc.SetInput("MessageBody", "")

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("[%s]", result)
}
