# DynamoDB Query
This activity provides your Flogo app the ability to execute a query against an Amazon DynamoDB

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
        "name": "DynamoDBKeyConditionExpression",
        "type": "string"
      },
      {
        "name": "DynamoDBExpressionAttributes",
        "type": "any"
      },
      {
        "name": "DynamoDBFilterExpression",
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
| DynamoDBKeyConditionExpression | The expression to search for (for example `Artist = :name` to search for results where the key `Artist` has the value of `:name`) |
| DynamoDBExpressionAttributes   | A JSON array representation of your expression attributes (using the example above that would be `[{"Name":":name", "Value":"Leon"}]` to search where the `Artist` is called `Leon`) |  
| DynamoDBFilterExpression       | The filter expression you want to apply on the result set before it is sent back to activity |

## Ouputs
| Output       | Description                                         |
|:-------------|:----------------------------------------------------|
| result       | The JSON representation of the result of your query |
| scannedCount | The number of items evaluated                       |