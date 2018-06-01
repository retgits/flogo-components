// Package dynamodbinsert inserts an object into Amazon DynamoDB
package dynamodbinsert

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	ivAwsAccessKeyID     = "awsAccessKeyID"
	ivAwsSecretAccessKey = "awsSecretAccessKey"
	ivAwsRegion          = "awsDefaultRegion"
	ivDynamoDBTableName  = "dynamoDBTableName"
	ivDynamoDBRecord     = "dynamoDBRecord"

	ovResult = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-dynamodbinsert")

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

// RecordAttribute is a structure representing the JSON payload for the record syntax
type RecordAttribute struct {
	Name  string
	Value string
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the inputs
	awsRegion := context.GetInput(ivAwsRegion).(string)
	dynamoDBTableName := context.GetInput(ivDynamoDBTableName).(string)
	dynamoDBRecord := context.GetInput(ivDynamoDBRecord)

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

	// Create a new login to the DynamoDB service
	dynamoService := dynamodb.New(awsSession)

	// Construct the expression attributes from the JSON payload
	var recordAttributes []RecordAttribute
	json.Unmarshal([]byte(dynamoDBRecord.(string)), &recordAttributes)

	recordAttributeMap := make(map[string]*dynamodb.AttributeValue)
	for _, attribute := range recordAttributes {
		recordAttributeMap[attribute.Name] = &dynamodb.AttributeValue{S: aws.String(attribute.Value)}
	}

	// Construct the DynamoDB Input
	input := &dynamodb.PutItemInput{
		TableName: aws.String(dynamoDBTableName),
		Item:      recordAttributeMap,
	}

	// Put the item in DynamoDB
	_, err1 := dynamoService.PutItem(input)
	if err1 != nil {
		log.Errorf("Error while executing query [%s]", err)
	} else {
		context.SetOutput(ovResult, "Added record to DynamoDB")
	}

	// Complete the activity
	return true, nil
}
