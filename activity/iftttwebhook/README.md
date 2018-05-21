# IFTTT WebHook
This activity provides your Flogo app the ability to trigger an IFTTT WebHook and use the full library of IFTTT to do cool things!

## Installation

```bash
flogo install github.com/retgits/flogo-components/activity/iftttwebhook
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/activity/iftttwebhook
```

## Schema
Inputs and Outputs:

```json
{
    "inputs": [
        {
            "name": "key",
            "type": "string",
            "required": true
        },
        {
            "name": "event",
            "type": "string",
            "required": true
        },
        {
            "name": "value1",
            "type": "string"
        },
        {
            "name": "value2",
            "type": "string"
        },
        {
            "name": "value3",
            "type": "string"
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
| Input  | Description                                                                                                          |
|:-------|:---------------------------------------------------------------------------------------------------------------------|
| key    | The key to connect to IFTTT (comes from the [Webhooks settings](https://ifttt.com/services/maker_webhooks/settings)). The `key` is the last part of the URL that is shown. For example, if the URL is `https://maker.ifttt.com/use/blablabla` the `Key` would be `blablabla` |
| event  | The Event Name you have used in your IFTTT recipe                                                                    |
| value1 | The first value you want to send                                                                                     |
| value2 | The second value you want to send                                                                                    |
| value3 | The third value you want to send                                                                                     |

_This activity will let you trigger IFTTT WebHooks with up to three parameters (this is a limit set by IFTTT)

## Ouputs
| Output      | Description                          |
|:------------|:-------------------------------------|
| result      | The HTTP status code of your request |