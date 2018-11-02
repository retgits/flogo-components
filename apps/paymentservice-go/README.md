# Payment Service
This sample Flogo application is used to demonstrate some key Flogo constructs:

- The Golang API for Flogo
- Input/Output mappings
- Complex object mapping

## Files
```bash
.
├── Dockerfile          <-- A Dockerfile to build a container based on an Alpine base image
├── main.go             <-- The Go source code for the app
├── Makefile            <-- A Makefile to help build and deploy the app
├── payment-go-svc.yml  <-- The Kubernetes deployment file
├── README.md           <-- This file
└── swagger.json        <-- The OpenAPI specification for the app
```

## Make targets
The [Makefile](./Makefile) has a few targets:
* **deps**: Get all the Go dependencies for the app
* **clean**: Remove the `dist` folder for a new deployment
* **clean-kube**: Remove all the deployed artifacts from Kubernetes
* **build-app**: Build an executable (and store it in the dist folder)
* **build-docker**: Build a Docker container from the contents of the `dist` folder
* **run-docker**: Run the Docker image with default settings
* **run-kube**: Deploy the app to Kubernetes

_For more detailed information on the commands that are executed you can check out the [Makefile](./Makefile)_

## API
After starting the app, it will register with two endpoints:
* **/api/expected-date/:invoiceId**: Generate a random payment date for an invoice
* **/swagger**: Get the OpenAPI specification for this app

## Visuals
If you want to see the visual representation of the app, check out the [paymentservice](../paymentservice) app
