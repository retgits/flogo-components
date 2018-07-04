package githubissues

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
    "name": "githubissues",
    "type": "flogo:activity",
    "ref": "github.com/retgits/flogo-components/activity/githubissues",
    "version": "0.2.0",
    "title": "GitHub Issues",
    "description": "Get the GitHub issues assigned to an authenticated user",
    "author": "retgits",
    "homepage": "https://github.com/retgits/flogo-components/tree/master/activity/githubissues",
    "inputs": [
        {
            "name": "token",
            "type": "string",
            "required": true
        },
        {
            "name": "timeInterval",
            "type": "integer",
            "required": true
        }
    ],
    "outputs": [
        {
            "name": "result",
            "type": "array"
        }
    ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
