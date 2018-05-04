package dynamodbquery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	tc.SetInput("AWSAccessKeyID", "")
	tc.SetInput("AWSSecretAccessKey", "")
	tc.SetInput("AWSDefaultRegion", "")
	tc.SetInput("DynamoDBTableName", "")
	tc.SetInput("DynamoDBKeyConditionExpression", "itemtype = :itemtype")
	//You can comment out the FilterExpression if you don't want to use it
	//tc.SetInput("DynamoDBFilterExpression", "<>")

	// You can pass in a string...
	b := `[{"Name":":itemtype","Value":"user"}]`

	// Or an object...
	//b := make(map[string]interface{})
	//b["Name"] = ":itemtype"
	//b["Value"] = "user"

	tc.SetInput("DynamoDBExpressionAttributes", b)
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	scannedCount := tc.GetOutput("scannedCount")
	consumedCapacity := tc.GetOutput("consumedCapacity")
	fmt.Printf("The ScannedCount of the query was: [%s]\n", scannedCount)
	fmt.Printf("The ConsumedCapacity of the query was: [%v]\n", consumedCapacity)
	fmt.Println("The Result of the query was:")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(result)
}
