package main

import (
	"fmt"
	"taskManagmant/models"
)

func main() {

	db := make(map[string]*models.TeamMember)

	alice := &models.TeamMember{Name: "Alice"}
	boob := &models.TeamMember{Name: "Bob"}
	maria := &models.TeamMember{Name: "Maria"}

	db["Alice"] = alice
	db["Boob"] = boob
	db["Maria"] = maria

	alice.AddTask("taskManagmant")
	alice.AddTask("Change password")

	db["Maria"].AddTask("Do Homework")
	db["Maria"].UpdateTaskStatus("Do Homework", "Done")

	alice.UpdateTaskStatus("taskManagmant", "done")

	fmt.Println("showing all task in the team...: ")

	for _, member := range db {

		models.DisplayAllTasks(member)
	}

	fmt.Println("showing all team members.... ")

	models.DisplayAllMembers(db)

}
