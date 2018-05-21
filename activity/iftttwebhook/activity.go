// Package iftttwebhook provides connectivity to IFTTT using the WebHooks service from IFTTT (https://ifttt.com/maker_webhooks)
package iftttwebhook

// Imports
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants
const (
	ivKey    = "key"
	ivEvent  = "event"
	ivValue1 = "value1"
	ivValue2 = "value2"
	ivValue3 = "value3"
	ovResult = "result"
)

// Payload is used to describe the payload to IFTTT with a maximum of
// three values (this limit is set by IFTTT)
type Payload struct {
	Value1 string `json:"value1,omitempty"`
	Value2 string `json:"value2,omitempty"`
	Value3 string `json:"value3,omitempty"`
}

// activityLog is the default logger for this class
var log = logger.GetLogger("activity-ifttt")

// IFTTTActivity describes the metadata of the activity as found in the activity.json file
type IFTTTActivity struct {
	metadata *activity.Metadata
}

// NewActivity will instantiate a new IFTTTActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &IFTTTActivity{metadata: metadata}
}

// Metadata will return the metadata of the IFTTTActivity
func (a *IFTTTActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval executes the activity and sends a message to IFTTT
func (a *IFTTTActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the connection details
	eventName := context.GetInput(ivEvent).(string)
	webhookKey := context.GetInput(ivKey).(string)

	// Build the IFTTT WebHook URL. To trigger the event it will make a POST request to the URL
	iftttURL := fmt.Sprintf("https://maker.ifttt.com/trigger/%s/with/key/%s", eventName, webhookKey)
	log.Infof("The WebHook URL is set to %s", iftttURL)

	// Create JSON payload. The data is completely optional, and you can also pass value1, value2, and value3 as query parameters or form variables.
	// This content will be passed on to the Action in your Recipe.
	payload := Payload{
		Value1: context.GetInput(ivValue1).(string),
		Value2: context.GetInput(ivValue2).(string),
		Value3: context.GetInput(ivValue3).(string),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return false, activity.NewError("Error occured while creating JSON Payload", "IFTTT-2001", nil)
	}
	body := bytes.NewReader(payloadBytes)

	// Send the POST message
	req, err := http.NewRequest("POST", iftttURL, body)
	if err != nil {
		return false, activity.NewError("Error occured sending message to IFTTT", "IFTTT-2002", nil)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, activity.NewError("Error occured receiving response from IFTTT", "IFTTT-2003", nil)
	}
	defer resp.Body.Close()

	log.Infof("WebHook returned with StatusCode %v", resp.StatusCode)

	// Set the return value
	context.SetOutput(ovResult, strconv.Itoa(resp.StatusCode))
	return true, nil
}
