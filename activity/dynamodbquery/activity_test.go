/*
Package dynamodbquery queries objects from Amazon DynamoDB


To be able to test this package you'll need to have access to Amazon DynamoDB. A sample table could be a table with the
name *data* and with *itemtype* as the partition key, and *itemid* as the sort key (both could be strings). Some sample
data (which can be generated with Mockaroo) can be

```
{
  "firstname": "John",
  "itemid": "57a98d98e4b00679b4a830af",
  "itemtype": "user",
  "lastname": "Doe",
  "password": "fec51acb3365747fc61247da5e249674cf8463c2",
  "username": "Jon_Doe"
}

{
  "firstname": "User",
  "itemid": "57a98d98e4b00679b4a830b2",
  "itemtype": "user",
  "lastname": "Name",
  "password": "e2de7202bb2201842d041f6de201b10438369fb8",
  "username": "user"
}

{
  "firstname": "Admin",
  "itemid": "57a98d98e4b00679b4a830b5",
  "itemtype": "admin",
  "lastname": "Name1",
  "password": "8f31df4dcc25694aeb0c212118ae37bbd6e47bcd",
  "username": "admin"
}
```

With this data you can test the KeyConditionExpression *itemtype = user* and a more complex
KeyConditionExpression *itemtype = user and itemid = 57a98d98e4b00679b4a830af*. The former will
return 2 objects, the latter only the John Doe record.
*/
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

// Update these variables before testing to match your own AWS account
const (
	awsAccessKeyID     = ""
	awsSecretAccessKey = ""
	awsDefaultRegion   = ""
	dynamoDBTableName  = "data"
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

// Test for a single condition string with no filtering
func TestEvalSingleString(t *testing.T) {

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
	tc.SetInput("DynamoDBTableName", dynamoDBTableName)
	tc.SetInput("DynamoDBKeyConditionExpression", "itemtype = :itemtype")

	// Prepare the Key Condition Expression as Name/Value pairs
	a := `[{"Name":":itemtype","Value":"user"}]`

	// Execute the activity
	tc.SetInput("DynamoDBExpressionAttributes", a)
	act.Eval(tc)

	// Check the result
	printOutput(tc)
}

// Test for a multiple condition string with no filtering
func TestEvalMultiString(t *testing.T) {

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
	tc.SetInput("DynamoDBTableName", dynamoDBTableName)
	tc.SetInput("DynamoDBKeyConditionExpression", "itemtype = :itemtype and itemid = :itemid")

	// Prepare the Key Condition Expression as Name/Value pairs
	a := `[{"Name":":itemtype","Value":"user"},{"Name":":itemid","Value":"57a98d98e4b00679b4a830b2"}]`

	// Execute the activity
	tc.SetInput("DynamoDBExpressionAttributes", a)
	act.Eval(tc)

	// Check the result
	printOutput(tc)
}

// Test for a single condition as JSON object with no filtering
func TestEvalJSONObject(t *testing.T) {

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
	tc.SetInput("DynamoDBTableName", dynamoDBTableName)
	tc.SetInput("DynamoDBKeyConditionExpression", "itemtype = :itemtype")

	// Prepare the Key Condition Expression as Name/Value pairs map
	a := make(map[string]interface{})
	a["Name"] = ":itemtype"
	a["Value"] = "user"

	// Execute the activity
	tc.SetInput("DynamoDBExpressionAttributes", a)
	act.Eval(tc)

	// Check the result
	printOutput(tc)
}

// Test for a multiple condition as JSON array with no filtering
func TestEvalJSONArray(t *testing.T) {

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
	tc.SetInput("DynamoDBTableName", dynamoDBTableName)
	tc.SetInput("DynamoDBKeyConditionExpression", "itemtype = :itemtype and itemid = :itemid")

	// Prepare the Key Condition Expression as Name/Value pairs maps
	a := make([]interface{}, 2)

	b := make(map[string]interface{})
	b["Name"] = ":itemtype"
	b["Value"] = "user"
	a[0] = b

	b = make(map[string]interface{})
	b["Name"] = ":itemid"
	b["Value"] = "57a98d98e4b00679b4a830b2"
	a[1] = b

	// Execute the activity
	tc.SetInput("DynamoDBExpressionAttributes", a)
	act.Eval(tc)

	// Check the result
	printOutput(tc)
}

func printOutput(tc *test.TestActivityContext) {
	result := tc.GetOutput("result")
	scannedCount := tc.GetOutput("scannedCount")
	consumedCapacity := tc.GetOutput("consumedCapacity")
	fmt.Printf("The ScannedCount of the query was: [%s]\n", scannedCount)
	fmt.Printf("The ConsumedCapacity of the query was: [%v]\n", consumedCapacity)
	fmt.Println("The Result of the query was:")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(result)
	fmt.Println("---")
}
