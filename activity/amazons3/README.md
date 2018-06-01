# Amazon S3

Activities for interacting with Amazon Simple Storage Service (S3).


## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/amazons3
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/amazons3
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
                "download",
                "upload",
                "delete"
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
            "name": "s3BucketName",
            "type": "string",
            "required": true
        },
        {
            "name": "s3Location",
            "type": "string",
            "required": true
        },
        {
            "name": "localLocation",
            "type": "string"
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
| Input              | Description    |
|:-------------------|:---------------|
| action             | The action you want to take, either `download`, `upload`, or `delete` |
| awsAccessKeyID     | Your AWS Access Key (only needed if you don't give your Lambda function rights to invoke S3 reads and puts) |
| awsSecretAccessKey | Your AWS Secret Key (only needed if you don't give your Lambda function rights to invoke S3 reads and puts) |
| awsRegion          | The AWS region your S3 bucket is in |
| s3BucketName       | The name of your S3 bucket |
| s3Location         | The file location on S3, this should be a full path (like `/bla/temp.txt`) |
| localLocation      | The `localLocation` is the full path to a file (like `/bla/temp.txt`) when uploading a file or the full path to a directory (like `./tmp`) when downloading a file |

## Ouputs
| Output    | Description    |
|:----------|:---------------|
| result    | The result will contain OK if the action was carried out successfully or will contain an error message |