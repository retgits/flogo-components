// Package gzip implements activities for reading and writing of gzip format compressed files, as specified in RFC 1952.
package gzip

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivAction          = "action"
	ivRemoveFile      = "removeFile"
	ivSourceFile      = "sourceFile"
	ivTargetDirectory = "targetDirectory"
	ovResult          = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-gzip")

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
	removeFile := context.GetInput(ivRemoveFile).(bool)
	sourceFile := context.GetInput(ivSourceFile).(string)
	targetDirectory := context.GetInput(ivTargetDirectory).(string)

	// See which action needs to be taken
	switch action {
	case "gunzip":
		// Extract file to disk
		err := unpackFileToDisk(sourceFile, targetDirectory, removeFile)
		if err != nil {
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, "OK")
		return true, nil
	case "gzip":
		// Extract file to disk
		err := gzipFile(sourceFile, targetDirectory, removeFile)
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

// Function to unzip the gzipped file
func unpackFileToDisk(sourceFile string, targetDirectory string, removeFile bool) error {
	// Open a file pointer to the gzipped file
	reader, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Read the gzipped file
	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	// We need the FNAME field (Header.Name) to reconstruct the file
	// if there is no Name present, we'll manually construct one
	if len(archive.Header.Name) == 0 {
		tempArray := strings.Split(sourceFile, "/")
		archive.Header.Name = strings.Replace(tempArray[len(tempArray)-1], ".gz", "", -1)
	}

	// Print the archive header
	log.Debugf("[File]   : [%s]", sourceFile)
	log.Debugf("[Comment]: [%s]", archive.Header.Comment)
	log.Debugf("[Extra]  : [%s]", archive.Header.Extra)
	log.Debugf("[ModTime]: [%s]", archive.Header.ModTime)
	log.Debugf("[Name]   : [%s]", archive.Header.Name)
	log.Debugf("[OS]     : [%s]", archive.Header.OS)

	// Create a new file pointer for the unpacked file
	writer, err := os.Create(filepath.Join(targetDirectory, archive.Name))
	if err != nil {
		return err
	}
	defer writer.Close()

	// Write the content of the gzip to the file
	_, err = io.Copy(writer, archive)

	// Play nice and remove the file
	if removeFile {
		err = os.Remove(sourceFile)
		if err != nil {
			return err
		}
	}

	return err
}

// Function to gzip the file
func gzipFile(sourceFile string, targetDirectory string, removeFile bool) error {
	// Open a file pointer to the file
	reader, err := os.Open(sourceFile)
	if err != nil {
		return err
	}

	// Set parameters for the gzipped file
	filename := filepath.Base(sourceFile)
	target := filepath.Join(targetDirectory, fmt.Sprintf("%s.gz", filename))

	// Create a new file pointer for the gzipped file
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	// Prepare the gzip writer
	archiver := gzip.NewWriter(writer)
	archiver.Name = filename
	defer archiver.Close()

	// Write the content to the gzip
	_, err = io.Copy(archiver, reader)

	// Play nice and remove the file
	if removeFile {
		err = os.Remove(sourceFile)
		if err != nil {
			return err
		}
	}

	return err
}
