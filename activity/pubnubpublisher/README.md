# PubNub Publisher

Publish messages to PubNub

## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/pubnubpublisher
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/pubnubpublisher
```

## Schema
Inputs and Outputs:

```json
{
    "inputs": [
        {
            "name": "publishKey",
            "type": "string",
            "required": true
        },
        {
            "name": "subscribeKey",
            "type": "string",
            "required": true
        },
        {
            "name": "uuid",
            "type": "string",
            "required": false
        },
        {
            "name": "channel",
            "type": "string",
            "required": true
        },
        {
            "name": "message",
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
| Input        | Description                                  |
|:-------------|:---------------------------------------------|
| publishKey   | The Publish Key from your PubNub Key Set     |
| subscribeKey | The Subscribe Key from your PubNub Key Set   |
| uuid         | The UUID for the Client connection           |
| channel      | The channel you want to publish a message to |
| message      | The actual message you want to send          |

## Ouputs
| Output    | Description    |
|:----------|:---------------|
| result    | The result will contain the UUID of the message that was published |