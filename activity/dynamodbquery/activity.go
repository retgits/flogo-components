// Package dynamodbquery queries objects from Amazon DynamoDB
package dynamodbquery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	ivAwsAccessKeyID                 = "awsAccessKeyID"
	ivAwsSecretAccessKey             = "awsSecretAccessKey"
	ivAwsRegion                      = "awsRegion"
	ivDynamoDBTableName              = "dynamoDBTableName"
	ivDynamoDBKeyConditionExpression = "dynamoDBKeyConditionExpression"
	ivDynamoDBExpressionAttributes   = "dynamoDBExpressionAttributes"
	ivDynamoDBFilterExpression       = "dynamoDBFilterExpression"

	ovResult           = "result"
	ovScannedCount     = "scannedCount"
	ovConsumedCapacity = "consumedCapacity"
)

// log is the default package logger
var log = logger.GetLogger("activity-dynamodbquery")

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
	dynamoDBTableName := context.GetInput(ivDynamoDBTableName).(string)
	dynamoDBKeyConditionExpression := context.GetInput(ivDynamoDBKeyConditionExpression).(string)
	dynamoDBExpressionAttributes := context.GetInput(ivDynamoDBExpressionAttributes)
	dynamoDBFilterExpression := context.GetInput(ivDynamoDBFilterExpression).(string)

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

	// Construct the expression attributes
	var expressionAttributes []ExpressionAttribute

	v := reflect.ValueOf(dynamoDBExpressionAttributes)
	switch v.Kind() {
	case reflect.String:
		json.Unmarshal([]byte(dynamoDBExpressionAttributes.(string)), &expressionAttributes)
	case reflect.Slice:
		fmt.Printf("slice")
	case reflect.Map:
		fmt.Printf("map")
	default:
		log.Errorf("Unknown type [%s]", reflect.TypeOf(dynamoDBExpressionAttributes).String())
		return true, fmt.Errorf("Unknown type [%s]", reflect.TypeOf(dynamoDBExpressionAttributes).String())
	}

	// Construct the expression attributes
	if reflect.TypeOf(dynamoDBExpressionAttributes).Kind() == reflect.Map {
		expressionAttributes = buildExpressionAttributesArray(dynamoDBExpressionAttributes.(map[string]interface{}))

	} else if reflect.TypeOf(dynamoDBExpressionAttributes).Kind() == reflect.Slice {
		tempArray := dynamoDBExpressionAttributes.([]interface{})
		for _, element := range tempArray {
			expressionAttributes = append(expressionAttributes, buildExpressionAttributesArray(element.(map[string]interface{}))...)
		}
	}

	log.Infof("%v", expressionAttributes)

	expressionAttributeMap := make(map[string]*dynamodb.AttributeValue)
	for _, attribute := range expressionAttributes {
		expressionAttributeMap[attribute.Name] = &dynamodb.AttributeValue{S: aws.String(attribute.Value)}
	}

	// Construct the DynamoDB query
	var queryInput = &dynamodb.QueryInput{}
	if dynamoDBFilterExpression == "" {
		queryInput = &dynamodb.QueryInput{
			TableName:                 aws.String(dynamoDBTableName),
			KeyConditionExpression:    aws.String(dynamoDBKeyConditionExpression),
			ExpressionAttributeValues: expressionAttributeMap,
			ReturnConsumedCapacity:    aws.String("TOTAL"),
		}
	} else {
		queryInput = &dynamodb.QueryInput{
			TableName:                 aws.String(dynamoDBTableName),
			KeyConditionExpression:    aws.String(dynamoDBKeyConditionExpression),
			ExpressionAttributeValues: expressionAttributeMap,
			FilterExpression:          aws.String(dynamoDBFilterExpression),
			ReturnConsumedCapacity:    aws.String("TOTAL"),
		}
	}

	// Prepare and execute the DynamoDB query
	var queryOutput, err1 = dynamoService.Query(queryInput)
	if err1 != nil {
		log.Errorf("Error while executing query [%s]", err1)
	} else {
		result := make([]map[string]interface{}, len(queryOutput.Items))
		// Loop over the result items and build a new map structure from it
		for index, element := range queryOutput.Items {
			dat := make(map[string]interface{})
			for key, value := range element {
				if value.N != nil {
					actual := *value.N
					dat[key] = actual
				}
				if value.S != nil {
					actual := *value.S
					dat[key] = actual
				}
			}
			result[index] = dat
		}
		// Set the output value in the context
		sc := *queryOutput.ScannedCount
		context.SetOutput(ovScannedCount, sc)
		cc := *queryOutput.ConsumedCapacity.CapacityUnits
		context.SetOutput(ovConsumedCapacity, cc)

		// Create a JSON representation from the result
		jsonString, _ := json.Marshal(result)
		var resultinterface interface{}
		d := json.NewDecoder(bytes.NewReader(jsonString))
		d.UseNumber()
		err = d.Decode(&resultinterface)
		f := map[string]interface{}{"results": resultinterface}
		context.SetOutput(ovResult, f)
	}
	// Complete the activity
	return true, nil
}

func buildExpressionAttributesArray(attribs map[string]interface{}) []ExpressionAttribute {
	var expressionAttributes []ExpressionAttribute
	attribValues := make([]string, 0, len(attribs))
	for _, v := range attribs {
		log.Infof("----[%s]", v.(string))
		attribValues = append(attribValues, v.(string))
	}
	for i := 0; i < len(attribValues); {
		expressionAttributes = append(expressionAttributes, ExpressionAttribute{Name: attribValues[i], Value: attribValues[i+1]})
		i += 2
	}
	return expressionAttributes
}
