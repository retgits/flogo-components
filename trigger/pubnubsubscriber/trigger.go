// Package pubnubsubscriber implements a trigger to receive messages from PubNub.
package pubnubsubscriber

import (
	"context"
	"fmt"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	pubnub "github.com/pubnub/go"
)

// log is the default package logger
var log = logger.GetLogger("trigger-pubnub")

// PubNubTrigger struct
type PubNubTrigger struct {
	metadata   *trigger.Metadata
	config     *trigger.Config
	pnInstance *pubnub.PubNub
	handlerMap map[string]*trigger.Handler
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &PubNubFactory{metadata: md}
}

// PubNubFactory PubNub Trigger factory
type PubNubFactory struct {
	metadata *trigger.Metadata
}

// New Creates a new trigger instance for a given id
func (t *PubNubFactory) New(config *trigger.Config) trigger.Trigger {
	return &PubNubTrigger{metadata: t.metadata, config: config}
}

// Metadata implements trigger.Trigger.Metadata
func (t *PubNubTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Initialize makes sure the trigger is completely set up for the Start method to work
func (t *PubNubTrigger) Initialize(ctx trigger.InitContext) error {
	// Make sure there is at least one settings item
	if t.config.Settings == nil {
		return fmt.Errorf("no settings found for trigger '%s'", t.config.Id)
	}

	// Make sure the publishKey item exists
	if _, ok := t.config.Settings["publishKey"]; !ok {
		return fmt.Errorf("no publishKey found for trigger '%s' in settings", t.config.Id)
	}

	// Make sure the subscribeKey item exists
	if _, ok := t.config.Settings["subscribeKey"]; !ok {
		return fmt.Errorf("no subscribeKey found for trigger '%s' in settings", t.config.Id)
	}

	// Create a new map to hold channels and handlers
	handlerMap := make(map[string]*trigger.Handler)

	// Init handlers
	for _, handler := range ctx.GetHandlers() {

		// Get the channel name
		channel := handler.GetStringSetting("channel")

		// Add the channel/handler combination to the map
		handlerMap[channel] = handler
	}

	t.handlerMap = handlerMap

	return nil
}

func (t *PubNubTrigger) Start() error {
	// Create a new PubNub config
	pnConfig := pubnub.NewConfig()
	pnConfig.PublishKey = t.config.GetSetting("publishKey")
	pnConfig.SubscribeKey = t.config.GetSetting("subscribeKey")

	// If the uuid item exists set the UUID of the client
	if _, ok := t.config.Settings["uuid"]; ok {
		pnConfig.UUID = t.config.GetSetting("uuid")
	}

	// Create a new PubNub instance and store that in the trigger struct
	pn := pubnub.NewPubNub(pnConfig)

	// Create a new Listener
	listener := pubnub.NewListener()
	pn.AddListener(listener)

	// Subscribe to all channels
	channelArray := make([]string, 0)
	for channel := range t.handlerMap {
		channelArray = append(channelArray, channel)
	}
	pn.Subscribe().Channels(channelArray).Execute()

	go func() {
		for {
			select {
			case status := <-listener.Status:
				switch status.Category {
				case pubnub.PNDisconnectedCategory:
					logger.Info("Received status [pubnub.PNDisconnectedCategory], this is the expected category for an unsubscribe. This means there was no error in unsubscribing from everything")
				case pubnub.PNConnectedCategory:
					logger.Info("Received status [pubnub.PNConnectedCategory], this is expected for a subscribe, this means there is no error or issue whatsoever")
				case pubnub.PNReconnectedCategory:
					logger.Info("Received status [pubnub.PNReconnectedCategory], this usually occurs if subscribe temporarily fails but reconnects. This means there was an error but there is no longer any issue")
				case pubnub.PNAccessDeniedCategory:
					logger.Info("Received status [pubnub.PNAccessDeniedCategory], this means that PAM does allow this client to subscribe to this channel and channel group configuration. This is another explicit error")
				}
			case message := <-listener.Message:
				logger.Debugf("%v", message)
				// Find the handler that is associated with the channel, else ignore the message
				if _, ok := t.handlerMap[message.Channel]; ok {
					onMessage(message, t.handlerMap[message.Channel])
				}
			case <-listener.Presence:
				// TODO: Precense allows you to subscribe to realtime Presence events, such as join, leave, and timeout, by UUID. This is currently not implemented
			}
		}
	}()

	return nil
}

func (t *PubNubTrigger) Stop() error {
	// Unsubscribe to all channels
	channelArray := make([]string, 0)
	for channel := range t.handlerMap {
		channelArray = append(channelArray, channel)
	}
	t.pnInstance.Unsubscribe().Channels(channelArray).Execute()

	return nil
}

func onMessage(message *pubnub.PNMessage, handler *trigger.Handler) {
	// Create a map to hold the trigger data
	triggerData := map[string]interface{}{
		"message":      message.Message,
		"channel":      message.Channel,
		"subscription": message.Subscription,
		"publisher":    message.Publisher,
		"timeToken":    strconv.FormatInt(message.Timetoken, 10),
	}

	// Execute the flow
	_, err := handler.Handle(context.Background(), triggerData)
	if err != nil {
		log.Infof("PubNub Error: %s", err.Error())
	}
}
