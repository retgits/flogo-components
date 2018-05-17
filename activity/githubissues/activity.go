// Package githubissues implements activities to get GitHub issues
package githubissues

// Imports
import (
	"bytes"
	ctx "context"
	"encoding/json"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Constants
const (
	ivGithubAccessToken = "token"
	ivTimeInterval      = "timeInterval"
	ovResult            = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-githubissues")

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
	// Get the data
	githubToken := context.GetInput(ivGithubAccessToken).(string)
	timeInterval := context.GetInput(ivTimeInterval).(int)

	// Create a new GitHub client
	ctxt := ctx.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctxt, ts)
	client := github.NewClient(tc)

	// Create a new time to make sure we check for new issues from the previous
	// execution of this function.
	interval := time.Duration(timeInterval) * time.Minute
	t := time.Now().Add(-interval)
	log.Infof("Check GitHub issues for the current user since [%s]", t)

	// Get all the issues assigned to the current authenticated user
	issueOpts := github.IssueListOptions{Since: t}
	issues, _, err := client.Issues.List(ctxt, false, &issueOpts)

	if err != nil {
		log.Error(err.Error())
		return true, err
	}

	log.Infof("GitHub returned %v issues", len(issues))
	result := make([]interface{}, len(issues))

	for idx, issue := range issues {
		result[idx] = issue
	}

	// Create a JSON representation from the result
	jsonString, _ := json.Marshal(result)
	var resultinterface interface{}
	d := json.NewDecoder(bytes.NewReader(jsonString))
	d.UseNumber()
	err = d.Decode(&resultinterface)
	datamap := map[string]interface{}{"results": resultinterface}

	// Set the output value in the context
	context.SetOutput(ovResult, datamap)

	return true, nil
}
