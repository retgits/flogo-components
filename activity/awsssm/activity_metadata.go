package awsssm

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
    "name": "amazonssm",
    "type": "flogo:activity",
    "ref": "github.com/retgits/flogo-components/activity/awsssm",
    "version": "0.3.0",
    "title": "AWS SSM",
    "description": "Store and Retrieve parameters from the Parameter Store in Amazon Simple Systems Manager (SSM)",
    "author": "retgits",
    "homepage": "https://github.com/retgits/flogo-components/tree/master/activity/awsssm",
    "inputs": [
        {
            "name": "action",
            "type": "string",
            "allowed": [
                "store",
                "retrieve",
                "retrieveList"
            ],
            "required": true
        },
        {
            "name": "awsAccessKeyID",
            "type": "string",
            "required": false
        },
        {
            "name": "awsSecretAccessKey",
            "type": "string",
            "required": false
        },
        {
            "name": "awsRegion",
            "type": "string",
            "required": true
        },
        {
            "name": "parameterName",
            "type": "string",
            "required": true
        },
        {
            "name": "decryptParameter",
            "type": "boolean"
        },
        {
            "name": "parameterValue",
            "type": "string"
        },
        {
            "name": "overwriteExistingParameter",
            "type": "boolean"
        },
        {
            "name": "parameterType",
            "type": "string",
            "allowed": [
                "String",
                "SecureString"
            ]
        }
    ],
    "outputs": [
        {
            "name": "result",
            "type": "any"
        }
    ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
