// Package grpctrigger implements a trigger to receive messages over gRPC.
package grpctrigger

import (
	"errors"
	"fmt"
	"net"
	"reflect"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"google.golang.org/grpc"
)

// log is the default package logger
var log = logger.GetLogger("trigger-grpc")

// Create a new map to hold the mapping of Service/RPC calls to Flogo handlers
var handlerMap = make(map[string]*trigger.Handler)

// Server struct, needed to register the gRPC server
type server struct{}

// GRPCTrigger struct
type GRPCTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &GRPCFactory{metadata: md}
}

// GRPCFactory gRPC Trigger factory
type GRPCFactory struct {
	metadata *trigger.Metadata
}

// New Creates a new trigger instance for a given id
func (t *GRPCFactory) New(config *trigger.Config) trigger.Trigger {
	return &GRPCTrigger{metadata: t.metadata, config: config}
}

// Metadata implements trigger.Trigger.Metadata
func (t *GRPCTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Initialize makes sure the trigger is completely set up for the Start method to work
func (t *GRPCTrigger) Initialize(ctx trigger.InitContext) error {
	// Init handlers
	for _, handler := range ctx.GetHandlers() {

		// Get the service name
		service := handler.GetStringSetting("service")

		// Get the rpc method name
		rpc := handler.GetStringSetting("rpc")

		// Add the channel/handler combination to the map
		handlerMap[fmt.Sprintf("%s-%s", service, rpc)] = handler
	}

	return nil
}

func (t *GRPCTrigger) Start() error {
	// Get the TCP port to listen on
	tcpPort := t.config.GetSetting("tcpPort")

	log.Infof("Starting gRPC TCP server on port %s", tcpPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", tcpPort))
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	registerServerMethods(s)

	if err := s.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (t *GRPCTrigger) Stop() error {
	return nil
}

// fillStruct maps the fields from a map[string]interface{} to a struct
func fillStruct(m map[string]interface{}, s interface{}) error {
	structValue := reflect.ValueOf(s).Elem()

	for name, value := range m {
		structFieldValue := structValue.FieldByName(name)

		if !structFieldValue.IsValid() {
			return fmt.Errorf("No such field: %s in obj", name)
		}

		if !structFieldValue.CanSet() {
			return fmt.Errorf("Cannot set %s field value", name)
		}

		val := reflect.ValueOf(value)
		if structFieldValue.Type() != val.Type() {
			return errors.New("Provided value type didn't match obj field type")
		}

		structFieldValue.Set(val)
	}
	return nil
}
