# gRPC trigger

Subscribe to gRPC messages

## Installation

```bash
flogo install github.com/retgits/flogo-components/trigger/grpctrigger
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/trigger/grpctrigger
```

## Schema
Inputs and Outputs:

```json
{
    "settings": [
        {
            "name": "tcpPort",
            "type": "string",
            "required": true
        },
        {
            "name": "protofileLocation",
            "type": "string",
            "required": true
        }
    ],
    "output": [
        {
            "name": "message",
            "type": "any"
        }
    ],
    "reply": [
        {
            "name": "data",
            "type": "any"
        }
    ],
    "handler": {
        "settings": [
            {
                "name": "service",
                "type": "string",
                "required": true
            },
            {
                "name": "rpc",
                "type": "string",
                "required": true
            }
        ]
    }
}
```
## Settings (generic settings for the trigger)
| Input             | Description                                  |
|:------------------|:---------------------------------------------|
| tcpPort           | The tcp port for the Flogo app to listen on  |
| protofileLocation | The location of the `.proto` file to use     |

## Ouputs (data sent to the flow)
| Output       | Description                                                  |
|:-------------|:-------------------------------------------------------------|
| message      | The JSON representation of the gRPC message                  |

## Reply (data sent back to the trigger)
| Input | Description                                  |
|:------|:---------------------------------------------|
| data  | The JSON representation of the gRPC message  |

## Handler (flow specific settings)
| Handler   | Description                                   |
|:----------|:----------------------------------------------|
| channel   | The channel you want to listen to             |

## Usage
Coming soon...