# Random String
This activity provides your Flogo app the ability to generate a random string consisting with the length you specify


## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/randomstring
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/randomstring
```

## Schema
Inputs and Outputs:

```json
{
"inputs":[
    {
      "name": "length",
      "type": "integer"
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
| Input   | Description    |
|:----------|:---------------|
| length | The length of the random string |

## Ouputs
| Output   | Description    |
|:----------|:---------------|
| result    | The random string |