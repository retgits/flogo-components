# EnvKey

Get environment variables or the provided fallback value

## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/envkey
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/envkey
```

## Schema
Inputs and Outputs:

```json
{
    "inputs":[
      {
        "name": "envkey",
        "type": "string",
        "required": true
      },
      {
        "name": "fallback",
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
| Input    | Description                                                       |
|:---------|:------------------------------------------------------------------|
| envkey   | The environment variable you want to get                          |
| fallback | The fallback value you want to provide if the variable is not set |

## Ouputs
| Output           | Description                                           |
|:-----------------|:------------------------------------------------------|
| result           | The value of the environment variable or the fallback |