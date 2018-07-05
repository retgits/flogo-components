# PubNub Subscriber

Subscribe to PubNub messages

## Installation

```bash
flogo install github.com/retgits/flogo-components/trigger/pubnub
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/trigger/pubnub
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
| channel      | The channel you want to publish a message to |
| message      | The actual message you want to send          |

## Ouputs
| Output    | Description                                   |
|:----------|:----------------------------------------------|
| message   | The message received from PubNub              |
| channel   | The channel on which the message was received |

## Handler
| Handler   | Description                                   |
|:----------|:----------------------------------------------|
| channel   | The channel you want to listen to             |