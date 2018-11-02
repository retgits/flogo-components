package openfaas

import (
	"fmt"
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

var jsonMetadata = `{
    "name": "openfaas",
    "type": "flogo:trigger",
    "shim": "plugin",
    "ref": "github.com/retgits/flogo-components/trigger/openfaas",
    "version": "0.0.1",
    "title": "OpenFaaS Trigger",
    "description": "OpenFaaS Trigger used to start a flow as a function.",
    "homepage": "https://github.com/retgits/flogo-contrib/tree/master/trigger/openfaas",
    "settings": [
    ],
    "output": [
      {
        "name": "context",
        "type": "object"
      },
      {
        "name": "evt",
        "type": "object"
      }
    ],
    "reply": [
      {
        "name": "data",
        "type": "any"
      },
      {
        "name": "status",
        "type": "integer",
        "value": 200
      }
    ]
  }`

// init create & register trigger factory
func init() {
	md := trigger.NewMetadata(jsonMetadata)
	// Adding this to make sure it works with the OpenFaaS build process
	// The way to check it actually runs in OpenFaaS is to see if the variable
	// fprocess is set to "./handler", which is done in the OpenFaaS docker
	// build process
	if os.Getenv("fprocess") == "./handler" {
		md.ID = fmt.Sprintf("handler/function/vendor/%s", md.ID)
	}
	factory := NewFactory(md)
	trigger.RegisterFactory(md.ID, factory)
}
