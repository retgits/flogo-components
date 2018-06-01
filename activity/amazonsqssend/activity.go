// Package amazonsqssend sends a message using Amazon SQS
package amazonsqssend

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	ivAwsAccessKeyID     = "awsAccessKeyID"
	ivAwsSecretAccessKey = "awsSecretAccessKey"
	ivAwsRegion          = "awsRegion"
	ivQueueURL           = "queueUrl"
	ivMessageBody        = "messageBody"
	ovResult             = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-amazonsqssend")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// ExpressionAttribute is a structure representing the JSON payload for the expression syntax
type ExpressionAttribute struct {
	Name  string
	Value string
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the inputs
	awsRegion := context.GetInput(ivAwsRegion).(string)
	queueURL := context.GetInput(ivQueueURL).(string)
	messageBody := context.GetInput(ivMessageBody).(string)

	// AWS Credentials, only if needed
	var awsAccessKeyID, awsSecretAccessKey = "", ""
	if context.GetInput(ivAwsAccessKeyID) != nil {
		awsAccessKeyID = context.GetInput(ivAwsAccessKeyID).(string)
	}
	if context.GetInput(ivAwsSecretAccessKey) != nil {
		awsSecretAccessKey = context.GetInput(ivAwsSecretAccessKey).(string)
	}

	// Create a session with Credentials only if they are set
	var awsSession *session.Session
	if awsAccessKeyID != "" && awsSecretAccessKey != "" {
		// Create new credentials using the accessKey and secretKey
		awsCredentials := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")

		// Create a new session with AWS credentials
		awsSession = session.Must(session.NewSession(&aws.Config{
			Credentials: awsCredentials,
			Region:      aws.String(awsRegion),
		}))
	} else {
		// Create a new session without AWS credentials
		awsSession = session.Must(session.NewSession(&aws.Config{
			Region: aws.String(awsRegion),
		}))
	}

	// Create a new login to the SQS service
	sqsService := sqs.New(awsSession)

	// Create a new SQS message
	sendMessageInput := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(messageBody),
	}

	//Send message to SQS
	response, err := sqsService.SendMessage(sendMessageInput)
	if err != nil {
		return true, err
	}

	//Set Message ID in the output
	context.SetOutput(ovResult, *response.MessageId)
	return true, nil
}
