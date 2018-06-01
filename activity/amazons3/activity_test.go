// Package amazons3 uploads or downloads files from Amazon Simple Storage Service (S3)
package amazons3

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEvalDownload(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("action", "download")
	tc.SetInput("awsAccessKeyID", "")
	tc.SetInput("awsSecretAccessKey", "")
	tc.SetInput("awsRegion", "")
	tc.SetInput("s3BucketName", "")
	tc.SetInput("s3Location", "")
	tc.SetInput("localLocation", "")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("[%s]", result)
}

func TestEvalUpload(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("action", "upload")
	tc.SetInput("awsAccessKeyID", "")
	tc.SetInput("awsSecretAccessKey", "")
	tc.SetInput("awsRegion", "")
	tc.SetInput("s3BucketName", "")
	tc.SetInput("s3Location", "")
	tc.SetInput("localLocation", "")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("[%s]", result)
}

func TestEvalDelete(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("action", "delete")
	tc.SetInput("awsAccessKeyID", "")
	tc.SetInput("awsSecretAccessKey", "")
	tc.SetInput("awsRegion", "")
	tc.SetInput("s3BucketName", "")
	tc.SetInput("s3Location", "")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("[%s]", result)
}
