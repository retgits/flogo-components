// Package stoplight implements an activity that looks at the incoming data elements
package stoplight

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

const (
	whitelistArray = `{ "one": [ { "name": "Packston", "id": "54" }, { "name": "Thia", "id": "88" }, { "name": "Woodie", "id": "78" }, { "name": "Charlotta", "id": "86" }, { "name": "Brinn", "id": "74" } ], "two": [ { "name": "Georgeanne", "id": "68" }, { "name": "Melicent", "id": "52" } ] }`
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

func TestEvalGreen(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("whitelistArray", whitelistArray)
	tc.SetInput("whitelist", "one")
	tc.SetInput("key", "54")
	tc.SetInput("value", "Packston")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	assert.Equal(t, result, "GREEN")
}

func TestEvalRed(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("whitelistArray", whitelistArray)
	tc.SetInput("whitelist", "one")
	tc.SetInput("key", "12")
	tc.SetInput("value", "Retgits")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	assert.Equal(t, result, "RED")
}

func TestEvalError(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("whitelistArray", whitelistArray)
	tc.SetInput("whitelist", "bla")
	tc.SetInput("key", "1337")
	tc.SetInput("value", "leet")
	_, err := act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	assert.Equal(t, result, "")
	assert.Error(t, err)
}
