# httptrigger

This trigger provides your flogo application the ability to start a flow via HTTP

_If your input is a JSON object you might want to use the [REST trigger](https://github.com/TIBCOSoftware/flogo-contrib/tree/master/trigger/rest). This trigger is specifically for cases where you receive form data (like `fizz=buzz`)_

## Installation
```bash
flogo install github.com/retgits/flogo-components/trigger/httptrigger
```
Link for flogo web:
```
https://github.com/retgits/flogo-components/trigger/httptrigger
```

## Schema
Settings, Outputs and Endpoint:

```json
{
  "settings": [
    {
      "name": "port",
      "type": "integer"
    }
  ],
  "output": [
    {
      "name": "pathParams",
      "type": "params"
    },
    {
      "name": "queryParams",
      "type": "params"
    },
    {
      "name": "header",
      "type": "params"
    },
    {
      "name": "content",
      "type": "object"
    }
  ],
  "endpoint": {
    "settings": [
      {
        "name": "method",
        "type": "string",
        "required" : true
      },
      {
        "name": "path",
        "type": "string",
        "required" : true
      }
    ]
  }
}
```
## Settings
### Trigger:
| Setting     | Description           |
|:------------|:----------------------|
| port        | The port to listen on |    

### Endpoint:
| Setting     | Description        |
|:------------|:-------------------|
| method      | The HTTP method    |         
| path        | The resource path  |


## Example Configurations

Triggers are configured via the triggers.json of your application. The following are some example configuration of the REST Trigger.

### POST
Configure the Trigger to handle a POST on /device

```json
{
  "triggers": [
    {
      "name": "flogo-rest",
      "settings": {
        "port": "8080"
      },
      "endpoints": [
        {
          "actionType": "flow",
          "actionURI": "embedded://new_device_flow",
          "settings": {
            "method": "POST",
            "path": "/device"
          }
        }
      ]
    }
  ]
}
```

### GET
Configure the Trigger to handle a GET on /device/:id

```json
{
  "triggers": [
    {
      "name": "flogo-rest",
      "settings": {
        "port": "8080"
      },
      "endpoints": [
        {
          "actionType": "flow",
          "actionURI": "embedded://get_device_flow",
          "settings": {
            "method": "GET",
            "path": "/device/:id"
          }
        }
      ]
    }
  ]
}
```
