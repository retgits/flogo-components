# Amazon Simple Email Service

Sends emails using Amazon Simple Email Service (SES)


## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/amazonses
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/amazonses
```

## Schema
Inputs and Outputs:

```json
{
    "inputs": [
        {
            "name": "to",
            "type": "string",
            "required": true
        },
        {
            "name": "from",
            "type": "string",
            "required": true
        },
        {
            "name": "content",
            "type": "string",
            "required": true
        },
        {
            "name": "subject",
            "type": "string",
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
| Input              | Description    |
|:-------------------|:---------------|
| to                 | The email address from which to send the address (must be a validated email address) |
| from               | The email address to which to send the email |
| content            | The content to send |
| subject            | The subject of the email |
| awsAccessKeyID     | Your AWS Access Key (only needed if you don't give your Lambda function rights to interact with Amazon SES) |
| awsSecretAccessKey | Your AWS Secret Key (only needed if you don't give your Lambda function rights to interact with Amazon SES) |
| awsRegion          | The AWS region from where you want to use SES (only needed if you don't give your Lambda function rights to interact with Amazon SES) |

## Ouputs
| Output    | Description    |
|:----------|:---------------|
| result    | The result from Amazon SES |
