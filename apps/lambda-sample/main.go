//go:generate go run ../../../TIBCOSoftware/flogo-lib/flogo/gen/gen.go -shim $GOPATH

package main

import (
	"context"
	"fmt"

	"github.com/TIBCOSoftware/flogo-contrib/activity/log"
	_ "github.com/TIBCOSoftware/flogo-contrib/activity/log"
	"github.com/TIBCOSoftware/flogo-contrib/trigger/lambda"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/flogo"
)

func shimApp() *flogo.App {

	app := flogo.NewApp()

	trg := app.NewTrigger(&lambda.LambdaTrigger{}, nil)
	trg.NewFuncHandler(nil, RunActivities)

	return app
}

func RunActivities(ctx context.Context, inputs map[string]*data.Attribute) (map[string]*data.Attribute, error) {

	fmt.Println(ctx)
	fmt.Println(inputs)

	in := map[string]interface{}{"message": "test"}
	_, err := flogo.EvalActivity(&log.LogActivity{}, in)

	if err != nil {
		return nil, err
	}

	response := make(map[string]interface{})

	response["id"] = "123"
	response["ref"] = "ref://blah"
	response["amount"] = "1"
	response["balance"] = "500"
	response["currency"] = "USD"

	ret := make(map[string]*data.Attribute)
	ret["code"], _ = data.NewAttribute("code", data.TypeInteger, 200)
	ret["data"], _ = data.NewAttribute("code", data.TypeAny, response)

	return ret, nil
}
