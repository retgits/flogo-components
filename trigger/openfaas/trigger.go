package openfaas

import (
	"context"
	"encoding/json"
	"flag"
	syslog "log"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("openfaas-trigger")
var singleton *OpenFaaSTrigger

// OpenFaaSTrigger OpenFaaS trigger struct
type OpenFaaSTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	handlers []*trigger.Handler
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &OpenFaaSFactory{metadata: md}
}

// OpenFaaSFactory OpenFaaS Trigger factory
type OpenFaaSFactory struct {
	metadata *trigger.Metadata
}

// New Creates a new trigger instance for a given id
func (t *OpenFaaSFactory) New(config *trigger.Config) trigger.Trigger {
	singleton = &OpenFaaSTrigger{metadata: t.metadata, config: config}
	return singleton

}

// Metadata implements trigger.Trigger.Metadata
func (t *OpenFaaSTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Initialize initializes the trigger
func (t *OpenFaaSTrigger) Initialize(ctx trigger.InitContext) error {
	t.handlers = ctx.GetHandlers()
	return nil
}

// Invoke starts the trigger and invokes the action registered in the handler
func Invoke() (map[string]interface{}, error) {
	log.Info("Starting OpenFaaS Trigger")
	syslog.Println("Starting OpenFaaS Trigger")

	// Parse the flags
	flag.Parse()

	// Looking up the arguments
	evtArg := flag.Lookup("evt")
	var evt interface{}

	// Unmarshall evt
	if err := json.Unmarshal([]byte(evtArg.Value.String()), &evt); err != nil {
		return nil, err
	}

	log.Debugf("Received evt: '%+v'\n", evt)
	syslog.Printf("Received evt: '%+v'\n", evt)

	// Get the context
	ctxArg := flag.Lookup("ctx")
	var funcCtx interface{}

	// Unmarshal ctx
	if err := json.Unmarshal([]byte(ctxArg.Value.String()), &funcCtx); err != nil {
		return nil, err
	}

	log.Debugf("Received ctx: '%+v'\n", funcCtx)
	syslog.Printf("Received ctx: '%+v'\n", funcCtx)

	//select handler, use 0th for now
	handler := singleton.handlers[0]

	inputData := map[string]interface{}{
		"context": funcCtx,
		"evt":     evt,
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
		log.Debugf("OpenFaaS Trigger Error: %s", err.Error())
		return nil, err
	}

	flowResponse := map[string]interface{}{
		"data":   replyData,
		"status": replyStatus,
	}

	return flowResponse, err

}

// Start implements util.Managed.Start
func (t *OpenFaaSTrigger) Start() error {
	return nil
}

// Stop implements util.Managed.Stop
func (t *OpenFaaSTrigger) Stop() error {
	return nil
}
