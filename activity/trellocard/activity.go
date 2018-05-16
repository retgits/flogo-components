// Package trellocard implements activities to create cards in Trello
package trellocard

// Imports
import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/adlio/trello"
)

// Constants
const (
	ivTrelloToken        = "token"
	ivTrelloKey          = "appkey"
	ivTrelloList         = "list"
	ivTrelloCardPosition = "position"
	ivCardTitle          = "title"
	ivCardDescription    = "description"
	ovResult             = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-trellocard")

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
	trelloToken := context.GetInput(ivTrelloToken).(string)
	trelloKey := context.GetInput(ivTrelloKey).(string)
	trelloList := context.GetInput(ivTrelloList).(string)
	cardTitle := context.GetInput(ivCardTitle).(string)
	trelloCardPosition := context.GetInput(ivTrelloCardPosition).(string)
	cardDescription := context.GetInput(ivCardDescription).(string)

	// Create a new Trello client
	trelloClient := trello.NewClient(trelloKey, trelloToken)

	// Create a Trello card struct
	card := trello.Card{
		Name:   cardTitle,
		Desc:   cardDescription,
		IDList: trelloList,
	}

	// Create the actual card in Trello
	err = trelloClient.CreateCard(&card, trello.Defaults())
	if err != nil {
		context.SetOutput(ovResult, err.Error())
		return true, err
	}

	// Move the card to the correct position
	switch trelloCardPosition {
	case "top":
		card.MoveToTopOfList()
	case "bottom":
		card.MoveToBottomOfList()
	}

	// Set the output value in the context
	context.SetOutput(ovResult, "OK")

	return true, nil
}
