package eventbus

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
)

var (
	ErrorParseMessageBody = errors.New("Error parse message body")
)

func HandleAddTaskToGoal(message IEventMessage) error {
	request := goal.EventAddTaskToGoal{}
	if err := json.Unmarshal(message.Body(), &request); err != nil {
		return ErrorParseMessageBody
	}
	fmt.Println("Add Task To Goal " + request.GoalID)
	return nil
}

func HandleCheckInTask(message IEventMessage) error {
	request := goal.EventCheckInTask{}
	if err := json.Unmarshal(message.Body(), &request); err != nil {
		return ErrorParseMessageBody
	}
	fmt.Println("Check In To Goal " + string(request.GoalID) + " with message " + request.Message)
	return nil
}
