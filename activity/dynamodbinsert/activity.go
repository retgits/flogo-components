// Package dynamodbinsert inserts a record in an Amazon DynamoDB
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
	ivAWSAccessKeyID     = "AWSAccessKeyID"
	ivAWSSecretAccessKey = "AWSSecretAccessKey"
	ivAWSDefaultRegion   = "AWSDefaultRegion"
	ivDynamoDBTableName  = "DynamoDBTableName"
	ivDynamoDBRecord     = "DynamoDBRecord"

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
	awsAccessKeyID := context.GetInput(ivAWSAccessKeyID).(string)
	awsSecretAccessKey := context.GetInput(ivAWSSecretAccessKey).(string)
	awsDefaultRegion := context.GetInput(ivAWSDefaultRegion).(string)
	dynamoDBTableName := context.GetInput(ivDynamoDBTableName).(string)
	dynamoDBRecord := context.GetInput(ivDynamoDBRecord)

	// Create new credentials using the accessKey and secretKey
	awsCredentials := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")

	// Create a new session to AWS
	awsSession := session.Must(session.NewSession(&aws.Config{
		Credentials: awsCredentials,
		Region:      aws.String(awsDefaultRegion),
	}))

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
