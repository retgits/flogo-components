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
To use this trigger you'll need to have gRPC and Protocol Buffers V3 installed. Because these dependencies are not installed in Flogo Web by default, you can only build an app using the CLI (designing an app will work perfectly in the Flogo Web UI)

_The below install instructions are from the [gRPC quickstart](https://grpc.io/docs/quickstart/go.html)_

### Install gRPC
To install gRPC use the following command
```bash
$ go get -u google.golang.org/grpc
```

### Install Protocol Buffers v3
Install the protoc compiler that is used to generate gRPC service code. The simplest way to do this is to download pre-compiled binaries for your platform(`protoc-<version>-<platform>.zip`) from here: https://github.com/google/protobuf/releases

* Unzip this file.
* Update the environment variable `PATH` to include the path to the protoc binary file.

Next, install the protoc plugin for Go
```bash
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

The compiler plugin, protoc-gen-go, will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler, protoc, to find it.
```bash
$ export PATH=$PATH:$GOPATH/bin
```

### Create a .proto file
The basis of gRPC is the .proto file. In order to successfully build the app you'll need to have one. An example proto file looks like:
```
syntax = "proto3";

option java_multiple_files = true;
option java_package = "com.leon.helloleon";
option java_outer_classname = "HelloLeonProto";

package helloleon;

service SaySomethingCool {
    rpc Awesomeness(InMessage) returns (OutMessage) {}
}

message InMessage {
    string name = 1;
}

message OutMessage {
    string name = 1;
    string awesomeness = 2;
}
```

### Install trigger
To install the trigger into your Web UI, create a new app and click on the “+” icon on the left hand side of the screen. From there click on “Install new” and paste `https://github.com/retgits/flogo-components/trigger/grpctrigger` into the input dialog to get this new trigger.

### Create an app
With everythong installed you can create the app using the Web UI as you would for any other app. Note that while constructing the `return` activity, you start each field with a capital letter (generating the structs made all fields exported and with Go that means they start with a capital letter). For the example above, a valid return would be
```json
{
  "Name": "You",
  "Awesomeness": "Incredible!"
}
```

### Build
To build the app, you'll have to export it first. From the main screen of your app (where all flows are listed), click "export" and select "App". You can use the following commands to build the app
```bash
# Create the Flogo app
$ flogo create -f <app.json> <appname>
# Build an executable
$ cd <appname>
$ flogo build -shim receive_grpc_message
# Start the app
$ ./bin/<appname>
```