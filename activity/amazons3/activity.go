// Package amazons3 uploads or downloads files from Amazon Simple Storage Service (S3)
package amazons3

import (
	"os"
	"path/filepath"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	ivAction             = "action"
	ivAwsAccessKeyID     = "awsAccessKeyID"
	ivAwsSecretAccessKey = "awsSecretAccessKey"
	ivAwsRegion          = "awsRegion"
	ivS3BucketName       = "s3BucketName"
	ivLocalLocation      = "localLocation"
	ivS3Location         = "s3Location"
	ovResult             = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-amazons3")

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

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the action
	action := context.GetInput(ivAction).(string)
	awsRegion := context.GetInput(ivAwsRegion).(string)
	s3BucketName := context.GetInput(ivS3BucketName).(string)
	// localLocation is a file when uploading a file or a directory when downloading a file
	localLocation := context.GetInput(ivLocalLocation).(string)
	s3Location := context.GetInput(ivS3Location).(string)

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

	// See which action needs to be taken
	switch action {
	case "download":
		err := downloadFileFromS3(awsSession, localLocation, s3Location, s3BucketName)
		if err != nil {
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, "OK")
		return true, nil
	case "upload":
		err := uploadFileToS3(awsSession, localLocation, s3Location, s3BucketName)
		if err != nil {
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, "OK")
		return true, nil
	case "delete":
		err := deleteFileFromS3(awsSession, s3Location, s3BucketName)
		if err != nil {
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, "OK")
		return true, nil
	}

	// Set the output value in the context
	context.SetOutput(ovResult, "NOK")

	return true, nil
}

// Function to download a file from an S3 bucket
func downloadFileFromS3(awsSession *session.Session, directory string, s3Location string, s3BucketName string) error {
	// Create an instance of the S3 Manager
	s3Downloader := s3manager.NewDownloader(awsSession)

	// Create a new temporary file
	f, err := os.Create(filepath.Join(directory, s3Location))
	if err != nil {
		return err
	}

	// Prepare the download
	objectInput := &s3.GetObjectInput{
		Bucket: aws.String(s3BucketName),
		Key:    aws.String(s3Location),
	}

	// Download the file to disk
	_, err = s3Downloader.Download(f, objectInput)
	if err != nil {
		return err
	}

	return nil
}

// Function to delete a file from an S3 bucket
func deleteFileFromS3(awsSession *session.Session, s3Location string, s3BucketName string) error {
	// Create an instance of the S3 Manager
	s3Session := s3.New(awsSession)

	objectDelete := &s3.DeleteObjectInput{
		Bucket: aws.String(s3BucketName),
		Key:    aws.String(s3Location),
	}

	// Delete the file from S3
	_, err := s3Session.DeleteObject(objectDelete)
	if err != nil {
		return err
	}

	return nil
}

// Function to upload a file from an S3 bucket
func uploadFileToS3(awsSession *session.Session, localFile string, s3Location string, s3BucketName string) error {
	// Create an instance of the S3 Manager
	s3Uploader := s3manager.NewUploader(awsSession)

	// Create a file pointer to the source
	reader, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Prepare the upload
	uploadInput := &s3manager.UploadInput{
		Bucket: aws.String(s3BucketName),
		Key:    aws.String(s3Location),
		Body:   reader,
	}

	// Upload the file
	_, err = s3Uploader.Upload(uploadInput)
	if err != nil {
		return err
	}

	return nil
}
