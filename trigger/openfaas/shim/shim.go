package function

import (
	"encoding/json"
	"flag"
	"fmt"

	fl "github.com/retgits/flogo-components/trigger/openfaas"
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
	result, err := fl.Invoke()
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
