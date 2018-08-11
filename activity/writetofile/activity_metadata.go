package writetofile

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
    "name": "writetofile",
    "type": "flogo:activity",
    "ref": "github.com/retgits/flogo-components/activity/writetofile",
    "version": "0.1.0",
    "title": "Write file",
    "description": "Write to a file",
    "author": "retgits",
    "homepage": "https://github.com/retgits/flogo-components/tree/master/activity/writetofile",
    "inputs": [
        {
            "name": "filename",
            "type": "string",
            "required": true
        },
        {
            "name": "content",
            "type": "string",
            "required": true
        },
        {
            "name": "append",
            "type": "bool"
        },
        {
            "name": "create",
            "type": "bool"
        }
    ],
    "outputs": [
        {
            "name": "result",
            "type": "string"
        }
    ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
