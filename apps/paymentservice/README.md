# Payment Service
This sample Flogo application is used to demonstrate some key Flogo constructs:

- Input/Output mappings
- Complex object mapping

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
export HTTPPORT=9999
./paymentservice
```

Test the application by opening your browser or via CURL and getting the following URL: http://localhost:9999/api/expected-date/1234

The following result should appear:

```json
{"expectedDate":"2018-02-20","id":"1234"}
```

## Importing into Flogo Web
If you want to import this flow into Flogo Web, please note there is a this is a current limitation in Flogo Web that it doesn't automatically import activities you do not have installed yet. To workaround this you can open an existing flow (or create a new flow) and import the activities using the "Install new activity" button. For this app you'll need to install:

* Random Number: https://github.com/retgits/flogo-components/activity/randomnumber
* Add to Date: https://github.com/retgits/flogo-components/activity/addtodate
