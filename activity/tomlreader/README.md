# Read and Query TOML
This activity provides your Flogo app the ability to read and query TOML files

## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/tomlreader
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/tomlreader
```

## Schema
Inputs and Outputs:

```json
{
"inputs": [
        {
            "name": "filename",
            "type": "string",
            "required": true
        },
        {
            "name": "key",
            "type": "string",
            "required": true
        }
    ],
    "outputs": [
        {
            "name": "result",
            "type": "array"
        }
    ]
}
```
## Inputs
| Input    | Description                                                                     |
|:---------|:--------------------------------------------------------------------------------|
| filename | The name of the file you want to write to (like `data.txt` or `./tmp/data.txt`) |
| key      | The key you want to search for in the TOML file                                 |

## Ouputs
| Output      | Description                                                        |
|:------------|:-------------------------------------------------------------------|
| result      | The array representing the result of searching the key             |