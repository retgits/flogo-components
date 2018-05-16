# Send SQS Message
This activity provides your Flogo app the ability to send a message over Amazon SQS


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
            "name": "AWSAccessKeyID",
            "type": "string",
            "required": true
        },
        {
            "name": "AWSSecretAccessKey",
            "type": "string",
            "required": true
        },
        {
            "name": "AWSDefaultRegion",
            "type": "string",
            "required": true
        },
        {
            "name": "QueueUrl",
            "type": "string",
            "required": true
        },
        {
            "name": "MessageBody",
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
| Input              | Description                              |
|:-------------------|:-----------------------------------------|
| AWSAccessKeyID     | Your AWS Access Key                      |
| AWSSecretAccessKey | Your AWS Secret Key (keep this secret!)  |
| AWSDefaultRegion   | The region your queue is in              |
| QueueUrl           | The URL of your SQS queue                |
| MessageBody        | The body of the message you want to send |


## Ouputs
| Output    | Description                       |
|:----------|:----------------------------------|
| result    | The ID of the message sent to SQS |
