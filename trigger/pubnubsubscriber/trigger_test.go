// Package pubnubsubscriber implements a trigger to receive messages from PubNub.
package pubnubsubscriber

import (
	"io/ioutil"
)

var jsonTestMetadata = getTestJsonMetadata()

func getTestJsonMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}
