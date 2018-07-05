# PubNub Subscriber

Subscribe to PubNub messages

## Installation

```bash
flogo install github.com/retgits/flogo-components/trigger/pubnubsubscriber
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/trigger/pubnubsubscriber
```

## Schema
Inputs and Outputs:

```json
{
    "settings":[
        {
          "name": "publishKey",
          "type": "string",
          "required" : true
        },
        {
          "name": "subscribeKey",
          "type": "string",
          "required" : true
        },
        {
          "name": "uuid",
          "type": "string",
          "required" : false
        }
      ],
      "output": [
        {
          "name": "message",
          "type": "string"
        },
        {
          "name": "channel",
          "type": "string"
        },
        {
          "name": "subscription",
          "type": "string"
        },
        {
          "name": "publisher",
          "type": "string"
        },
        {
          "name": "timeToken",
          "type": "string"
        }
      ],
      "handler": {
        "settings": [
          {
            "name": "channel",
            "type": "string",
            "required" : true
          }
        ]
      }
}
```
## Inputs
| Input        | Description                                  |
|:-------------|:---------------------------------------------|
| publishKey   | The Publish Key from your PubNub Key Set     |
| subscribeKey | The Subscribe Key from your PubNub Key Set   |
| uuid         | The UUID for the Client connection           |

## Ouputs
| Output       | Description                                                  |
|:-------------|:-------------------------------------------------------------|
| message      | The message sent on the PubNub channel                       |
| channel      | The channel on which the message was received                |
| subscription | The channel group or wildcard subscription match (if exists) |
| publisher    | The UUID of the publisher                                    |
| timeToken    | The Timetoken for the message                                |

## Handler
| Handler   | Description                                   |
|:----------|:----------------------------------------------|
| channel   | The channel you want to listen to             |