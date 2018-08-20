package grpctrigger

import (
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

var jsonMetadata = `{
    "name": "grpctrigger",
    "type": "flogo:trigger",
    "shim": "plugin",
    "ref": "github.com/retgits/flogo-components/trigger/grpctrigger",
    "version": "0.2.0",
    "title": "Receive gRPC messages",
    "description": "gRPC trigger",
    "author": "retgits",
    "homepage": "https://github.com/retgits/flogo-components/trigger/grpctrigger",
    "settings": [
        {
            "name": "tcpPort",
            "type": "string",
            "required": true
        },
        {
            "name": "protofileLocation",
            "type": "string",
            "required": true
        }
    ],
    "output": [
        {
            "name": "message",
            "type": "any"
        }
    ],
    "reply": [
        {
            "name": "data",
            "type": "any"
        }
    ],
    "handler": {
        "settings": [
            {
                "name": "service",
                "type": "string",
                "required": true
            },
            {
                "name": "rpc",
                "type": "string",
                "required": true
            }
        ]
    }
}`

// init create & register trigger factory
func init() {
	md := trigger.NewMetadata(jsonMetadata)
	trigger.RegisterFactory(md.ID, NewFactory(md))
}
