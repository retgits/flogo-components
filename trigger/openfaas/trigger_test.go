package openfaas

import (
	"io/ioutil"
)

var jsonTestMetadata = getTestJsonMetadata()

func getTestJsonMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

const testConfig string = `{
  "id": "flogo-rest",
  "ref": "github.com/TIBCOSoftware/flogo-contrib/trigger/lambda",
  "settings": {
  },
  "handlers": [
	{
	  "actionId": "my_test_flow",
	  "settings": {
	  }
	}
  ]
}
`
