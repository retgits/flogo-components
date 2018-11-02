# Invoice Service
This sample Flogo application is used to demonstrate some key Flogo constructs:

- Input/Output mappings
- Complex object mapping
- Invoke REST service

## Files
```bash
.
├── Dockerfile           <-- A Dockerfile to build a container based on an Alpine base image
├── Makefile             <-- A Makefile to help build and deploy the app
├── invoice-svc.yml      <-- The Kubernetes deployment file
├── invoiceservice.json  <-- The Flogo app
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
* **/api/invoices/:id**: Get the invoice details for the invoice ID. The app also reaches out to the [paymentservice](../paymentservice) app

## Building the app
To build via the Flogo CLI, simply download the invoiceservice.json to your local machine and create the app structure:

```{r, engine='bash', count_lines}
flogo create -f invoiceservice.json invoiceservice
cd invoiceservice
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
export PAYMENTSERVICE=http://localhost:9999/api/expected-date/:id
./invoiceservice
```

_Note that this app does require the paymentservice app to be running as well_

Test the application by opening your browser or via CURL and getting the following URL: http://localhost:9998/api/invoices/2345

The following result should appear:

```json
{"amount":1456,"balance":456,"currency":"USD","expectedPaymentDate":"2018-02-28","id":"2345","ref":"INV-2345"}
```

_Note that the expectedPaymentDate comes from the paymentservice app. If that app isn't running or if the URL is misconfigured it will return as null_

## Importing into Flogo Web
If you want to import this flow into Flogo Web, please note there is a this is a current limitation in Flogo Web that it doesn't automatically import activities you do not have installed yet. To workaround this you can open an existing flow (or create a new flow) and import the activities using the "Install new activity" button. For this app you'll need to install:

* Random Number: https://github.com/retgits/flogo-components/activity/randomnumber
* Random String: https://github.com/retgits/flogo-components/activity/randomstring
* Combine: https://github.com/jvanderl/flogo-components/activity/combine

## Nothing but Go?
If you want to see the same app, built with the Flogo Go API, check out the [invoiceservice-go](../invoiceservice-go) app