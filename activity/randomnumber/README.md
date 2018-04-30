# Random Number
This activity provides your Flogo app the ability to generate a random unique number between the min and max value


## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/randomnumber
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/randomnumber
```

## Schema
Inputs and Outputs:

```json
{
"inputs":[
    {
      "name": "min",
      "type": "integer"
    },
    {
      "name": "max",
      "type": "integer"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "integer"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| min | The minimum value of the random number |
| max | The maximum value of the random number |

## Ouputs
| Output   | Description    |
|:----------|:---------------|
| result    | The random number |