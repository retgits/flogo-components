# AWS Systems Manager

Store and Retrieve parameters from the Parameter Store in Amazon Simple Systems Manager (SSM)


## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/awsssm
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/awsssm
```

## Schema
Inputs and Outputs:

```json
{
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
            "required": true
        },
        {
            "name": "awsSecretAccessKey",
            "type": "string",
            "required": true
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
```
## Inputs
| Input                      | Description    |
|:---------------------------|:---------------|
| action                     | The action you want to take, either `store` (create a new parameter), `retrieve` (get a single parameter), or `retrieveList` (get multiple parameters) |
| awsAccessKeyID             | Your AWS Access Key (only needed if you don't give your Lambda function rights to interact with AWS SSM) |
| awsSecretAccessKey         | Your AWS Secret Key (only needed if you don't give your Lambda function rights to interact with AWS SSM) |
| parameterName              | The name of the parameter when the action is `store` or `retrieve` (like `param1`), or a comma separated list of parameters when the action is `retrieveList` (like `param1,param2`) |
| decryptParameter           | Describes whether the parameter should be decrypted if the action is `retrieve` or `retrieveList` |
| parameterValue             | The value of the parameter if the action is `store` (like `myAwesomeValue`) |
| overwriteExistingParameter | If the action is `store` this parameter describes whether to overwrite the value if the parameter already exists |
| parameterType              | The type of the parameter if the action is `store`, this can be either `String` (non-encrypted) or `SecureString` (encrypted with the default key of your account)

## Ouputs
| Output    | Description    |
|:----------|:---------------|
| result    | The result will contain a JSON representation of your result (like `{results:{"param1":"value1","param2":"value2"}}`) |
