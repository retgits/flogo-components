// Package pubnubpublisher implements activities to publish messages to PubNub.
package pubnubpublisher

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	pubnub "github.com/pubnub/go"
)

const (
	ivPublishKey   = "publishKey"
	ivSubscribeKey = "subscribeKey"
	ivChannel      = "channel"
	ivMessage      = "message"
	ivUUID         = "uuid"
	ovResult       = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-pubnubpublisher")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	publishKey := context.GetInput(ivPublishKey).(string)
	subscribeKey := context.GetInput(ivSubscribeKey).(string)
	channel := context.GetInput(ivChannel).(string)
	message := context.GetInput(ivMessage).(string)
	uuid := context.GetInput(ivUUID).(string)

	config := pubnub.NewConfig()
	config.PublishKey = publishKey
	config.SubscribeKey = subscribeKey
	if len(uuid) > 0 {
		config.UUID = uuid
	}

	pn := pubnub.NewPubNub(config)

	res, status, err := pn.Publish().Channel(channel).Message(message).Execute()

	// Set the output value in the context
	if err != nil {
		context.SetOutput(ovResult, status.Error.Error())
		return false, err
	}

	logger.Debugf("Timestamp of the publish response: [%v]", res.Timestamp)

	context.SetOutput(ovResult, status.UUID)
	return true, nil
}
