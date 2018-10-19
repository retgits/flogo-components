package specialtrigger

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("trigger")
var specialTrigger *SpecialTrigger

// SpecialTrigger struct
type SpecialTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	handlers []*trigger.Handler
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &SpecialFactory{metadata: md}
}

// SpecialFactory Special Trigger factory
type SpecialFactory struct {
	metadata *trigger.Metadata
}

// New Creates a new trigger instance for a given id
func (t *SpecialFactory) New(config *trigger.Config) trigger.Trigger {
	specialTrigger = &SpecialTrigger{metadata: t.metadata, config: config}
	return specialTrigger
}

// Metadata implements trigger.Trigger.Metadata
func (t *SpecialTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Initialize makes sure the trigger is completely set up for the Start method to work
func (t *SpecialTrigger) Initialize(ctx trigger.InitContext) error {
	// Init handlers
	t.handlers = ctx.GetHandlers()
	return nil
}

func (t *SpecialTrigger) Start() error {
	return nil
}

func (t *SpecialTrigger) Stop() error {
	return nil
}

// Invoke starts the trigger and invokes the action registered in the handler
func Invoke(r http.Request) (map[string]interface{}, error) {

	log.Info("Starting Special Trigger")
	syslog.Println("Starting Special Trigger")

	// Unmarshall evt
	var evt interface{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	if err := json.Unmarshal(buf.Bytes(), &evt); err != nil {
		return nil, err
	}

	handler := specialTrigger.handlers[0]

	inputData := map[string]interface{}{
		"message": evt,
	}

	results, err := handler.Handle(context.Background(), inputData)

	var replyData interface{}
	var replyStatus int

	if len(results) != 0 {
		dataAttr, ok := results["data"]
		if ok {
			replyData = dataAttr.Value()
		}
		code, ok := results["status"]
		if ok {
			replyStatus, _ = data.CoerceToInteger(code.Value())
		}
	}

	if err != nil {
		log.Debugf("Special Trigger Error: %s", err.Error())
		return nil, err
	}

	flowResponse := map[string]interface{}{
		"data":   replyData,
		"status": replyStatus,
	}
	return flowResponse, err
}
