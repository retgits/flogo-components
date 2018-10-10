// Package globalcache implements a wrapper around go-cache for Flogo. go-cache is an in-memory key:value
// store/cache similar to memcached that is suitable for applications running on a single machine.
// Its major advantage is that, being essentially a thread-safe map[string]interface{} with expiration
// times, it doesn't need to serialize or transmit its contents over the network.
package globalcache

import (
	"fmt"
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

func TestEvalGet(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "GET")
	tc.SetInput("key", "retgits")
	_, err := act.Eval(tc)

	// Get result attribute
	result := tc.GetOutput("result")

	// Verify result
	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestEvalDelete(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "DELETE")
	tc.SetInput("key", "retgits")
	_, err := act.Eval(tc)

	// Get result attribute
	result := tc.GetOutput("result")

	// Verify result
	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestEvalInitialize(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "INITIALIZE")
	tc.SetInput("expiryTime", 5)
	tc.SetInput("purgeTime", 10)
	_, err := act.Eval(tc)

	// Get result attribute
	result := tc.GetOutput("result")

	// Verify result
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestEvalSetAndInitialize(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "SET")
	tc.SetInput("key", "retgits")
	tc.SetInput("value", "Awesome!")
	_, err := act.Eval(tc)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get result attribute
	result := tc.GetOutput("result")

	// Verify result
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestEvalSetAndGet(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "SET")
	tc.SetInput("key", "retgits")
	tc.SetInput("value", "Awesome!")
	_, err := act.Eval(tc)

	// Get result attribute
	result := tc.GetOutput("result")
	fmt.Printf("%v\n", result)

	// Verify result
	assert.NoError(t, err)
	assert.Empty(t, result)

	tc = test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "GET")
	tc.SetInput("key", "retgits")
	_, err = act.Eval(tc)

	// Get result attribute
	result = tc.GetOutput("result")

	// Verify result
	assert.NoError(t, err)
	assert.Equal(t, result, "Awesome!")

	tc = test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "GET")
	tc.SetInput("key", "AnotherName")
	_, err = act.Eval(tc)

	// Get result attribute
	result = tc.GetOutput("result")

	// Verify result
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestEvalInitializeWithLoadset(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "INITIALIZE")
	tc.SetInput("expiryTime", 5)
	tc.SetInput("purgeTime", 10)
	tc.SetInput("loadset", `{"key":"value","retgits":"Awesome!"}`)
	_, err := act.Eval(tc)

	// Get result attribute
	result := tc.GetOutput("result")

	// Verify result
	assert.NoError(t, err)
	assert.Empty(t, result)

	tc = test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "GET")
	tc.SetInput("key", "retgits")
	_, err = act.Eval(tc)

	// Get result attribute
	result = tc.GetOutput("result")

	// Verify result
	assert.NoError(t, err)
	assert.Equal(t, result, "Awesome!")

	tc = test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "GET")
	tc.SetInput("key", "AnotherName")
	_, err = act.Eval(tc)

	// Get result attribute
	result = tc.GetOutput("result")

	// Verify result
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestEvalDeleteWithCache(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Setup attributes
	tc.SetInput("action", "DELETE")
	tc.SetInput("key", "retgits")
	_, err := act.Eval(tc)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get result attribute
	result := tc.GetOutput("result")

	// Verify result
	assert.NoError(t, err)
	assert.Empty(t, result)
}
