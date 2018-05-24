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
        },
        {
            "name": "filters",
            "type": "string"
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
| filters  | A filter string (see below)                                                     |

## Ouputs
| Output      | Description                                                        |
|:------------|:-------------------------------------------------------------------|
| result      | The array representing the result of searching the key             |

## filters
The current implementation of this activity supports two filtering mechanisms:

* Filter by value: Return the configuration item if one of the values in that item contains the value you're searching for
* Filter by key: Return the configuration item if one of the keys of that item equals the value you're searching for

### Filter by value
The filter `ValueContains(retgits)` would return all items where part any value in the item matches `retgits`

### Filter by key
The filter `KeyEquals(type,app)` would return all items which have a key called `type` and where the value of that key equals `app`

### Combining filters
Filters are applied sequentially and you can specify multiple filters separated by a `/`. The filter string `KeyEquals(type,app)/ValueContains(retgits)` would first perform a **Filter by key** and filter the result even further by applying the **Filter by value**