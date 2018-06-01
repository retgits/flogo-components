# Send SQS Message

Send a message using Amazon Simple Queue Service (SQS)

## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/amazonsqssend
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/amazonsqssend
```

## Schema
Inputs and Outputs:

```json
{
    "inputs": [
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
            "name": "queueUrl",
            "type": "string",
            "required": true
        },
        {
            "name": "messageBody",
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
```
## Inputs
| Input              | Description                                                                                                 |
|:-------------------|:------------------------------------------------------------------------------------------------------------|
| awsAccessKeyID     | Your AWS Access Key (only needed if you don't give your Lambda function rights to interact with Amazon SQS) |
| awsSecretAccessKey | Your AWS Secret Key (only needed if you don't give your Lambda function rights to interact with Amazon SQS) |
| awsRegion          | The region your queue is in                                                                                 |
| queueUrl           | The URL of your SQS queue                                                                                   |
| messageBody        | The body of the message you want to send                                                                    |


## Ouputs
| Output    | Description                       |
|:----------|:----------------------------------|
| result    | The ID of the message sent to SQS |
