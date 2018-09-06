package pubnubpublisher

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
    "name": "pubnubpublisher",
    "type": "flogo:activity",
    "ref": "github.com/retgits/flogo-components/activity/pubnubpublisher",
    "version": "0.1.0",
    "title": "PubNub Publisher",
    "description": "Publish messages to PubNub",
    "author": "retgits",
    "homepage": "https://github.com/retgits/flogo-components/tree/master/activity/pubnubpublisher",
    "inputs": [
        {
            "name": "publishKey",
            "type": "string",
            "required": true
        },
        {
            "name": "subscribeKey",
            "type": "string",
            "required": true
        },
        {
            "name": "uuid",
            "type": "string",
            "required": false
        },
        {
            "name": "channel",
            "type": "string",
            "required": true
        },
        {
            "name": "message",
            "type": "string",
            "required": true
        }
    ],
    "outputs": [
        {
            "name": "result",
            "type": "string"
        }
    ]
}`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
