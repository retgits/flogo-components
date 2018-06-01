# DynamoDB Query

Query objects from Amazon DynamoDB

## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/dynamodbquery
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/dynamodbquery
```

## Schema
Inputs and Outputs:

```json
{
"inputs":[
      {
        "name": "awsAccessKeyID",
        "type": "string"
      },
      {
        "name": "awsSecretAccessKey",
        "type": "string"
      },
      {
        "name": "awsRegion",
        "type": "string"
      },
      {
        "name": "dynamoDBTableName",
        "type": "string"
      },
      {
        "name": "dynamoDBKeyConditionExpression",
        "type": "string"
      },
      {
        "name": "dynamoDBExpressionAttributes",
        "type": "any"
      },
      {
        "name": "dynamoDBFilterExpression",
        "type": "string"
      }
    ],
    "outputs": [
    {
      "name": "result",
      "type": "any"
    },
    {
      "name": "scannedCount",
      "type": "string"
    },
    {
      "name": "consumedCapacity",
      "type": "double"
    }
  ]
}
```
## Inputs
| Input                          | Description    |
|:-------------------------------|:---------------|
| awsAccessKeyID                 | Your AWS Access Key (only needed if you don't give your Lambda function rights to interact with Amazon DyanmoDB)                       |
| awsSecretAccessKey             | Your AWS Secret Key (only needed if you don't give your Lambda function rights to interact with Amazon DyanmoDB)   |
| awsRegion                      | The AWS region you're running DynamoDB in |
| dynamoDBTableName              | The name of your DynamoDB table           |
| dynamoDBKeyConditionExpression | The expression to search for (for example `Artist = :name` to search for results where the key `Artist` has the value of `:name`) |
| dynamoDBExpressionAttributes   | A JSON array representation of your expression attributes (using the example above that would be `[{"Name":":name", "Value":"Leon"}]` to search where the `Artist` is called `Leon`) |  
| dynamoDBFilterExpression       | The filter expression you want to apply on the result set before it is sent back to activity |

## Ouputs
| Output           | Description                                         |
|:-----------------|:----------------------------------------------------|
| result           | The JSON representation of the result of your query |
| scannedCount     | The number of items evaluated                       |
| consumedCapacity | The number of capacity units used by the query      |