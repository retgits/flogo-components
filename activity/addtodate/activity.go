// Package addtodate provides your Flogo app the ability to add a specified number of units (days, months, or years) to a date
package addtodate

import (
	"time"

	"github.com/project-flogo/core/activity"
)

var activityMetadata = activity.ToMetadata(&Input{}, &Output{})

// Activity is an AddToDate Activity implementation
type Activity struct{}

// Metadata return the metadata for the activity
func (a *Activity) Metadata() *activity.Metadata {
	return activityMetadata
}

// New creates a new AddToDate activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	return &Activity{}, nil
}

// Eval executes the activity
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	// Get the inputs
	input := Input{}
	err = ctx.GetInputObject(&input)
	if err != nil {
		return false, err
	}

	d := time.Now()

	if input.Date != "" {
		d, _ = time.Parse("2006-01-02", input.Date)
	}

	switch input.Units {
	case "days":
		d = d.AddDate(0, 0, 1*input.Number)
	case "months":
		d = d.AddDate(0, 1*input.Number, 0)
	case "years":
		d = d.AddDate(1*input.Number, 0, 0)
	}

	// Set the output value in the context
	output := Output{
		Result: d.Format("2006-01-02"),
	}

	err = ctx.SetOutputObject(&output)
	if err != nil {
		return false, err
	}

	return true, nil
}
