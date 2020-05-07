package addtodate

import (
	"github.com/project-flogo/core/data/coerce"
)

// Input is the input into the AddToDate activity
type Input struct {
	// Date is the start date to add the units to
	Date string `md:"date"`

	// Number is the amount of units to add to the date
	Number int `md:"number"`

	// Units is the string representation of what needs to be added (days, months, or years)
	Units string `md:"number"`
}

// Output is the result of the AddToDate activity
type Output struct {
	// Result is the new date is ISO format (yyyy-mm-dd)
	Result string `md:"result"`
}

// FromMap converts the values from a map into the struct Input
func (i *Input) FromMap(values map[string]interface{}) error {
	date, err := coerce.ToString(values["date"])
	if err != nil {
		return err
	}
	i.Date = date

	number, err := coerce.ToInt(values["number"])
	if err != nil {
		return err
	}
	i.Number = number

	units, err := coerce.ToString(values["units"])
	if err != nil {
		return err
	}
	i.Units = units

	return nil
}

// ToMap converts the struct Input into a map
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"date":   i.Date,
		"number": i.Number,
		"units":  i.Units,
	}
}

// ToMap converts the struct Output into a map
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}

// FromMap converts the values from a map into the struct Output
func (o *Output) FromMap(values map[string]interface{}) error {
	var err error

	o.Result, err = coerce.ToString(values["result"])
	if err != nil {
		return err
	}

	return nil
}
