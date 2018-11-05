# OpenFaaS

> [OpenFaaS®](https://github.com/openfaas/faas) (Functions as a Service) is a framework for building Serverless functions with Docker and Kubernetes which has first-class support for metrics. Any process can be packaged as a function enabling you to consume a range of web events without repetitive boiler-plate coding.

## How it works

With OpenFaaS you can create your own Go-based functions and that means your Flogo apps (built with the Golang API) can be deployed there too! To make it work, OpenFaaS leverages a dockerfile that executes the build of the function and makes the docker image available to wherever you're running OpenFaaS.

## Schema
Settings, Outputs:

```json
{
  "settings": [
  ],
  "output": [
    {
      "name": "context",
      "type": "object"
    },
    {
      "name": "evt",
      "type": "object"
    }
  ]
}
```

## Dockerfile

Due to the build process of OpenFaaS and the required metadata for Project Flogo, the dockerfile to build these Go based apps needs to be modified slightly. The [dockerfile](./sampleapp/dockerfile) in the `sampleapp` folder has been modified to accomodate for those changes. On line 17, a small statement will update the metadata files with the correct `ID` for OpenFaaS.

## Sample app

As with the Lambda trigger, Flogo apps using the OpenFaaS trigger will break down into a few parts as you can see in the [handler.go](./sampleapp/code/handler.go) file

### Handle

This method takes care of starting and enabling the Flogo app to handle requests from the OpenFaaS environment

### init

This method takes care of the initialization of the Flogo app

### shimApp

This method takes care of creating the Flogo app and registering the OpenFaaS trigger with the Flogo engine

### RunActivities

RunActivities is where the magic happens. This is where you get the input from any event that might trigger your OpenFaaS function in a map called `evt`