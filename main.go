package main

import (
	"fmt"

	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

func main() {
	shiro := member.NewMember("Shiro")
	sampleGoal := goal.NewGoal("Goal 1", "Goal 1 Description", shiro)

	sampleGoal.AddTask("Task 1", "", 10, "time")
	sampleGoal.AddTask("Task 2", "", 50, "unit")

	sampleGoal.CheckIn("Task 1", 5, "This is a sample checkin")

	fmt.Println(sampleGoal.Progress())
}
