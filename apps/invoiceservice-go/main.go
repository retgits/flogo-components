//go:generate go run ../../../../TIBCOSoftware/flogo-lib/flogo/gen/gen.go $GOPATH
package main

import (
	"context"
	"fmt"
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
	app := flogo.NewApp()

	// Convert the HTTPPort to an integer
	port, err := strconv.Atoi(httpport)

	// Register the HTTP trigger
	trg := app.NewTrigger(&rt.RestTrigger{}, map[string]interface{}{"port": port})
	trg.NewFuncHandler(map[string]interface{}{"method": "GET", "path": "/api/invoices/:id"}, handler)

	e, err := flogo.NewEngine(app)

	if err != nil {
		logger.Error(err)
		return
	}

	engine.RunEngine(e)
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
	in = map[string]interface{}{"uri": fmt.Sprintf("%s%s", paymentservice, id)}
	out, err = flogo.EvalActivity(&rest.RESTActivity{}, in)
	if err != nil {
		return nil, err
	}

	// The return message is a map[string]*data.Attribute which we'll have to construct
	response := make(map[string]*data.Attribute)

	attr, _ := data.NewAttribute("id", data.TypeString, id)
	response["id"] = attr
	attr, _ = data.NewAttribute("ref", data.TypeString, ref)
	response["ref"] = attr
	attr, _ = data.NewAttribute("amount", data.TypeString, amount)
	response["amount"] = attr
	attr, _ = data.NewAttribute("balance", data.TypeString, balance)
	response["balance"] = attr
	//attr, _ = data.NewAttribute("expectedPaymentDate", data.TypeString, id)
	//response["expectedPaymentDate"] = attr
	attr, _ = data.NewAttribute("currency", data.TypeString, "USD")
	response["currency"] = attr

	return response, nil
}
