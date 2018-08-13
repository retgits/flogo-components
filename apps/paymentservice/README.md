# Payment Service
This sample Flogo application is used to demonstrate some key Flogo constructs:

- Input/Output mappings
- Complex object mapping

## Files
```bash
.
├── Dockerfile           <-- A Dockerfile to build a container based on an Alpine base image
├── Makefile             <-- A Makefile to help build and deploy the app
├── payment-svc.yml      <-- The Kubernetes deployment file
├── paymentservice.json  <-- The Flogo app
├── README.md            <-- This file
```

## Make targets
The [Makefile](./Makefile) has a few targets:
* **clean**: Remove the `dist` folder for a new deployment
* **clean-kube**: Remove all the deployed artifacts from Kubernetes
* **build-app**: Build an executable (both in Linux format as well as the native format of the current OS)
* **build-docker**: Build a Docker container
* **run-docker**: Run the Docker image with default settings
* **run-kube**: Deploy the app to Kubernetes

_For more detailed information on the commands that are executed you can check out the [Makefile](./Makefile)_

## API
After starting the app, it will register one endpoints:
* **/api/expected-date/:invoiceId**: Generate a random payment date for an invoice

## Building the app
To build via the Flogo CLI, simply download the paymentservice.json to your local machine and create the app structure:

```{r, engine='bash', count_lines}
flogo create -f paymentservice.json paymentservice
cd paymentservice
```

Now you can simply build the application and leverage the HTTP REST trigger as the entry point:

```{r, engine='bash', count_lines}
flogo build -e
```

The -e switch indicates that the binary should embed the flogo.json.

## Run the application

Now that the application has been built, run the application:

```{r, engine='bash', count_lines}
cd bin
export HTTPPORT=9998
./paymentservice
```

Test the application by opening your browser or via CURL and getting the following URL: http://localhost:9998/api/expected-date/1234

The following result should appear:

```json
{"expectedDate":"2018-02-20","id":"1234"}
```

## Importing into Flogo Web
If you want to import this flow into Flogo Web, please note there is a this is a current limitation in Flogo Web that it doesn't automatically import activities you do not have installed yet. To workaround this you can open an existing flow (or create a new flow) and import the activities using the "Install new activity" button. For this app you'll need to install:

* Random Number: https://github.com/retgits/flogo-components/activity/randomnumber
* Add to Date: https://github.com/retgits/flogo-components/activity/addtodate

## Nothing but Go?
If you want to see the same app, built with the Flogo Go API, check out the [paymentservice-go](../paymentservice-go) app