//go:generate go run $GOPATH/src/github.com/TIBCOSoftware/flogo-lib/flogo/gen/gen.go $GOPATH
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/TIBCOSoftware/flogo-contrib/activity/log"
	"github.com/TIBCOSoftware/flogo-contrib/activity/rest"
	rt "github.com/TIBCOSoftware/flogo-contrib/trigger/rest"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/engine"
	"github.com/TIBCOSoftware/flogo-lib/flogo"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/retgits/flogo-components/activity/randomnumber"
)

var (
	httpport       = os.Getenv("HTTPPORT")
	paymentservice = os.Getenv("PAYMENTSERVICE")
)

func main() {
	// Create a new Flogo app
	app := appBuilder()

	e, err := flogo.NewEngine(app)

	if err != nil {
		logger.Error(err)
		return
	}

	engine.RunEngine(e)
}

func appBuilder() *flogo.App {
	app := flogo.NewApp()

	// Convert the HTTPPort to an integer
	port, err := strconv.Atoi(httpport)
	if err != nil {
		logger.Error(err)
	}

	// Register the HTTP trigger
	trg := app.NewTrigger(&rt.RestTrigger{}, map[string]interface{}{"port": port})
	trg.NewFuncHandler(map[string]interface{}{"method": "GET", "path": "/api/invoices/:id"}, handler)
	trg.NewFuncHandler(map[string]interface{}{"method": "GET", "path": "/swaggerspec"}, SwaggerSpec)

	return app
}

// SwaggerSpec is the function that gets executedto retrieve the SwaggerSpec
func SwaggerSpec(ctx context.Context, inputs map[string]*data.Attribute) (map[string]*data.Attribute, error) {
	// The return message is a map[string]*data.Attribute which we'll have to construct
	response := make(map[string]interface{})
	ret := make(map[string]*data.Attribute)

	fileData, err := ioutil.ReadFile("swagger.json")
	if err != nil {
		ret["code"], _ = data.NewAttribute("code", data.TypeInteger, 500)
		response["msg"] = err.Error()
	} else {
		ret["code"], _ = data.NewAttribute("code", data.TypeInteger, 200)
		var data map[string]interface{}
		if err := json.Unmarshal(fileData, &data); err != nil {
			panic(err)
		}
		response = data
	}

	ret["data"], _ = data.NewAttribute("data", data.TypeAny, response)

	return ret, nil

}

func handler(ctx context.Context, inputs map[string]*data.Attribute) (map[string]*data.Attribute, error) {
	// Get the ID from the path
	id := inputs["pathParams"].Value().(map[string]string)["id"]

	// Execute the log activity
	in := map[string]interface{}{"message": id, "flowInfo": "true", "addToFlow": "true"}
	_, err := flogo.EvalActivity(&log.LogActivity{}, in)
	if err != nil {
		return nil, err
	}

	// Generate a random number for the amount
	// There are definitely better ways to do this with Go, but this keeps the flow consistent with the UI version
	in = map[string]interface{}{"min": 1000, "max": 2000}
	out, err := flogo.EvalActivity(&randomnumber.MyActivity{}, in)
	if err != nil {
		return nil, err
	}
	amount := strconv.Itoa(out["result"].Value().(int))

	// Instead of using the combine activity we'll concat the strings together
	ref := fmt.Sprintf("INV-%v", id)

	// Generate a random number for the balance
	// There are definitely better ways to do this with Go, but this keeps the flow consistent with the UI version
	in = map[string]interface{}{"min": 0, "max": 2000}
	out, err = flogo.EvalActivity(&randomnumber.MyActivity{}, in)
	if err != nil {
		return nil, err
	}
	balance := strconv.Itoa(out["result"].Value().(int))

	// Call out to another service
	in = map[string]interface{}{"method": "GET", "uri": fmt.Sprintf("%s%s", paymentservice, id)}
	logger.Info(in)
	out, err = flogo.EvalActivity(&rest.RESTActivity{}, in)
	if err != nil {
		return nil, err
	}
	expectedDate := out["result"].Value().(map[string]interface{})["expectedDate"].(string)

	// The return message is a map[string]*data.Attribute which we'll have to construct
	response := make(map[string]interface{})
	response["id"] = id
	response["ref"] = ref
	response["amount"] = amount
	response["balance"] = balance
	response["expectedPaymentDate"] = expectedDate
	response["currency"] = "USD"

	ret := make(map[string]*data.Attribute)
	ret["code"], _ = data.NewAttribute("code", data.TypeInteger, 200)
	ret["data"], _ = data.NewAttribute("data", data.TypeAny, response)

	return ret, nil
}
