# Stoplight

Performs a lookup of names and ids in a provided lookup list

## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/stoplight
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/stoplight
```

## Schema
Inputs and Outputs:

```json
{
    "inputs": [
        {
            "name": "whitelistArray",
            "type": "object",
            "required": true
        },
        {
            "name": "whitelist",
            "type": "string",
            "required": true
        },
        {
            "name": "key",
            "type": "string",
            "required": true
        },
        {
            "name": "value",
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
| Input           | Description    |
|:----------------|:---------------|
| whitelistArray  | An array of JSON objects to search in (like `{ "one": [ { "name": "Packston", "id": "54" }, { "name": "Thia", "id": "88" }, { "name": "Woodie", "id": "78" }, { "name": "Charlotta", "id": "86" }, { "name": "Brinn", "id": "74" } ], "two": [ { "name": "Georgeanne", "id": "68" }, { "name": "Melicent", "id": "52" } ] }`) |
| whitelist       | The named array you want to search in (like `one`) |
| key             | The value of the id field (like `54`) |
| value           | The value of the name field (like `Packston`) |

## Ouputs
| Output    | Description    |
|:----------|:---------------|
| result    | `RED` or `GREEN` depending if the key and value were found in the whitelist |