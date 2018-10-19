package specialtrigger

import (
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

var jsonMetadata = `{
    "name": "specialtrigger",
    "type": "flogo:trigger",
    "ref": "github.com/retgits/flogo-components/trigger/specialtrigger",
    "version": "0.2.0",
    "title": "Does something special",
    "description": "Special Trigger",
    "author": "retgits",
    "homepage": "https://github.com/retgits/flogo-components/trigger/specialtrigger",
    "settings": [],
    "output": [
        {
            "name": "message",
            "type": "any"
        }
    ],
    "handler": {},
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
	trigger.RegisterFactory(md.ID, NewFactory(md))
}
