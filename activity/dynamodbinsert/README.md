# DynamoDB Insert
This activity provides your Flogo app the ability to insert a record in an Amazon DynamoDB

## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/dynamodbinsert
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/dynamodbinsert
```

## Schema
Inputs and Outputs:

```json
"inputs":[
      {
        "name": "AWSAccessKeyID",
        "type": "string"
      },
      {
        "name": "AWSSecretAccessKey",
        "type": "string"
      },
      {
        "name": "AWSDefaultRegion",
        "type": "string"
      },
      {
        "name": "DynamoDBTableName",
        "type": "string"
      },
      {
        "name": "DynamoDBRecord",
        "type": "any"
      }
    ],
    "outputs": [
      {
        "name": "result",
        "type": "any"
      }
    ]
```
## Inputs
| Input                          | Description    |
|:-------------------------------|:---------------|
| AWSAccessKeyID                 | Your AWS Access Key                       |
| AWSSecretAccessKey             | Your AWS Secret Key (keep this secret!)   |
| AWSDefaultRegion               | The AWS region you're running DynamoDB in |
| DynamoDBTableName              | The name of your DynamoDB table           |
| DynamoDBRecord                 | A JSON array representation of your record attributes you want to add. They are name/value pairs so adding an Artist with name Leon would be `[{"Name":"Artist", "Value":"Leon"}]`. |  

## Ouputs
| Output       | Description                                            |
|:-------------|:-------------------------------------------------------|
| result       | A string denoting if the record was successfully added |