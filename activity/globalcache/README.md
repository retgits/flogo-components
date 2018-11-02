# GlobalCache

GlobalCache implements a wrapper around [go-cache](https://github.com/patrickmn/go-cache) for Flogo. 

> go-cache is an in-memory key:value store/cache similar to memcached that is suitable for applications running on a single machine. Its major advantage is that, being essentially a thread-safe map[string]interface{} with expiration times, it doesn't need to serialize or transmit its contents over the network.

## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/globalcache
```

Link for flogo web:

```bash
https://github.com/retgits/flogo-components/activity/globalcache
```

## Schema

Inputs and Outputs:

```json
{
    "inputs": [
        {
            "name": "action",
            "type": "string",
            "allowed": [
                "INITIALIZE",
                "SET",
                "GET",
                "DELETE"
            ],
            "required": true
        },
        {
            "name": "key",
            "type": "string",
            "required": false
        },
        {
            "name": "value",
            "type": "string",
            "required": false
        },
        {
            "name": "expiryTime",
            "type": "int",
            "required": false
        },
        {
            "name": "purgeTime",
            "type": "int",
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

| Input      | Description    |
|:-----------|:---------------|
| action     | The action to perform on the cache. Allowed values are `INITIALIZE` (to create a new GlobalCache object), `SET` (to add a new key/value pair to the cache), `GET` (to retrieve an object from the cache), and `DELETE` (to remove on object from the cache) |
| key        | The name of the key to look for (during a `GET` or `DELETE` operation) or to add to the cache (during a `SET` operation)
| value      | The value to add to the cache |
| expiryTime | The expiration time of items in the cache in minutes (defaults to `5`). Using `-1` as expiryTime means items will never expire |
| purgeTime  | The purge time of expired items in the cache in minutes (defaults to `10`) |
| loadset    | A JSON object representing key/value pairs to load while initializing the cache (like `{"key":"value","retgits":"Awesome!"}`)|

## Ouputs

| Output    | Description                                                                                           |
|:----------|:------------------------------------------------------------------------------------------------------|
| result    | The value of the item in the cache, if it exists (will return an empty element if no entry was found) |