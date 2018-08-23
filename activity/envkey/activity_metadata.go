package envkey

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
    "name": "envkey",
    "type": "flogo:activity",
    "ref": "github.com/retgits/flogo-components/activity/envkey",
    "version": "0.1.0",
    "title": "EnvKey",
    "description": "Get Environment Variable",
    "author": "retgits",
    "homepage": "https://github.com/retgits/flogo-components/tree/master/activity/envkey",
    "inputs":[
      {
        "name": "envkey",
        "type": "string",
        "required": true
      },
      {
        "name": "fallback",
        "type": "string",
        "required": true
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
