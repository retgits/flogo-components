//go:generate go run $GOPATH/src/github.com/TIBCOSoftware/flogo-lib/flogo/gen/gen.go $GOPATH
// Package function implements the main logic of the function
package function

// The ever important imports
import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/config"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/flogo"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/retgits/flogo-components/trigger/openfaas"
)

// Handle a serverless request
func Handle(req []byte) string {
	// Setup event argument and assume that the incoming byte array is a JSON payload
	evtFlag := flag.Lookup("evt")
	if evtFlag == nil {
		flag.String("evt", string(req), "OpenFaaS Environment Arguments")
	} else {
		flag.Set("evt", string(req))
	}

	// Setup context argument
	openfaasContext := map[string]interface{}{
		"environment": "OpenFaaS",
	}
	ctxJSON, err := json.Marshal(openfaasContext)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		panic(err.Error())
	}

	ctxFlag := flag.Lookup("ctx")
	if ctxFlag == nil {
		flag.String("ctx", string(ctxJSON), "OpenFaaS Context Arguments")
	} else {
		flag.Set("ctx", string(ctxJSON))
	}

	// Invoke the flogo OpenFaaS trigger and handle the event
	result, err := openfaas.Invoke()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		panic(err.Error())
	}

	responseRaw, err := json.Marshal(result["data"])
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		panic(err.Error())
	}

	return string(responseRaw)
}

// Init makes sure that everything is ready to go!
func init() {
	config.SetDefaultLogLevel("ERROR")
	logger.SetLogLevel(logger.ErrorLevel)

	app := shimApp()

	e, err := flogo.NewEngine(app)

	if err != nil {
		logger.Error(err)
		return
	}

	e.Init(true)
}

// shimApp is used to build a new Flogo app and register the OpenFaaS trigger with the engine.
// The shimapp is used by the shim, which triggers the engine every time an event comes into the contaoner.
func shimApp() *flogo.App {
	// Create a new Flogo app
	app := flogo.NewApp()

	// Register the trigger with the Flogo app
	trg := app.NewTrigger(&openfaas.OpenFaaSTrigger{}, nil)
	trg.NewFuncHandler(nil, RunActivities)

	// Return a pointer to the app
	return app
}

// RunActivities is where the magic happens. This is where you get the input from any event that might trigger
// your OpenFaaS function in a map called evt (which is part of the inputs). The below sample,
// will simply return "Go Serverless v1.x! Your function executed successfully!" as a response.
func RunActivities(ctx context.Context, inputs map[string]*data.Attribute) (map[string]*data.Attribute, error) {
	message := "Go Serverless v1.x! Your function executed successfully!"

	// The response from the OpenFaaS function is always in the form of a JSON message.
	// In this case we're creating a structure with a single element called
	// message
	responseData := make(map[string]interface{})
	responseData["message"] = message

	// Because we're sending the result back to the API Gateway, it will expect to have
	// both an HTTP result code (called code) and some response data. In this case we're
	// sending back the data object we created earlier
	response := make(map[string]*data.Attribute)
	response["code"], _ = data.NewAttribute("code", data.TypeInteger, 200)
	response["data"], _ = data.NewAttribute("data", data.TypeAny, responseData)

	return response, nil
}
