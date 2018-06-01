// Package dynamodbinsert inserts an object into Amazon DynamoDB
package dynamodbinsert

import (
	"encoding/json"
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

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	// To test this example, you can create a dynamodb with table name Music
	// where the key is called Artist
	tc.SetInput("AWSAccessKeyID", "<Your Access Key>")
	tc.SetInput("AWSSecretAccessKey", "<Your Secret Key>")
	tc.SetInput("AWSDefaultRegion", "<Your region>")
	tc.SetInput("DynamoDBTableName", "Music")

	payload := []RecordAttribute{
		RecordAttribute{
			Name:  "Artist",
			Value: "Leon",
		},
		RecordAttribute{
			Name:  "SongTitle",
			Value: "I cant really sing",
		},
		RecordAttribute{
			Name:  "Message",
			Value: "Please dont sing",
		},
		RecordAttribute{
			Name:  "LastPostedBy",
			Value: "Leon",
		},
	}

	b, _ := json.Marshal(payload)

	tc.SetInput("DynamoDBRecord", string(b))
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("The Result of the query was:\n[%s]\n", result)
}
