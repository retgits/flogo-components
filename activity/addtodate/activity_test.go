// Package addtodate adds a specified number of units to a date.
package addtodate

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})

	assert.NotNil(t, ref)
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act, err := New(nil)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	input := &Input{
		Date:   "2018-02-17",
		Units:  "months",
		Number: 2,
	}

	tc.SetInputObject(input)

	_, err = act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("result")
	assert.Equal(t, result, "2018-04-17")

	input = &Input{
		Date:   "2018-02-17",
		Units:  "days",
		Number: 2,
	}

	tc.SetInputObject(input)

	_, err = act.Eval(tc)
	assert.Nil(t, err)

	result = tc.GetOutput("result")
	assert.Equal(t, result, "2018-02-19")

	input = &Input{
		Date:   "2018-02-17",
		Units:  "years",
		Number: 2,
	}

	tc.SetInputObject(input)

	_, err = act.Eval(tc)
	assert.Nil(t, err)

	result = tc.GetOutput("result")
	assert.Equal(t, result, "2020-02-17")
}
