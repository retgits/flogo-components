# Command Parser
This activity provides your Flogo app the ability to read and parse a commandline string

## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/commandparser
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/commandparser
```

## Schema
Inputs and Outputs:

```json
{
"inputs": [
        {
            "name": "commandString",
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
| Input         | Description              |
|:--------------|:-------------------------|
| commandString | The command line string  |

## Ouputs
| Output      | Description                            |
|:------------|:---------------------------------------|
| result      | A `map[string]string` of the arguments |

_The commandString `--type trigger --string "twitter" -dev` would result in a map `map[type:trigger string:twitter dev:true]`_