# Query Parser

Parses a query string into name/value pairs.


## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/queryparser
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/queryparser
```

## Schema
Inputs and Outputs:

```json
{
    "inputs": [
        {
            "name": "query",
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
| Input          | Description                        |
|:---------------|:-----------------------------------|
| query          | The query string you want to parse |

## Ouputs
| Output    | Description                                           |
|:----------|:------------------------------------------------------|
| result    | A `map[string]interface{}` containing the parsed data |