package notification

import (
	"fmt"

	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
)

type NotificationService struct {
}

func NewNotificationService() *NotificationService {
	service := NotificationService{}
	return &service
}

func (*NotificationService) NotifyAddTaskToGoal(event *goal.EventAddTaskToGoal) {
	fmt.Println("Sent email to user when new task was added to goal")
}

func (*NotificationService) NotifyCheckInTask(event *goal.EventCheckInTask) {
	fmt.Println("Sent email to user when task was checked in")
}
