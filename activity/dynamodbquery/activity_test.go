package dynamodbquery

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
	tc.SetInput("AWSAccessKeyID", "<<>>")
	tc.SetInput("AWSSecretAccessKey", "<<>>")
	tc.SetInput("AWSDefaultRegion", "<<>>")
	tc.SetInput("DynamoDBTableName", "<<>>")
	tc.SetInput("DynamoDBKeyConditionExpression", "itemtype = :itemtype")
	// You can comment out the FilterExpression if you don't want to use it
	tc.SetInput("DynamoDBFilterExpression", "<>")

	payload := []ExpressionAttribute{
		ExpressionAttribute{
			Name:  ":itemtype",
			Value: "user",
		},
	}

	b, _ := json.Marshal(payload)

	tc.SetInput("DynamoDBExpressionAttributes", string(b))
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	scannedCount := tc.GetOutput("scannedCount")
	fmt.Printf("The ScannedCount of the query was: [%s]\n", scannedCount)
	fmt.Printf("The Result of the query was:\n[%s]\n", result)
}
