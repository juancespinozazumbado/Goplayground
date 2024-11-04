package main

import (
	"fmt"
	"taskManagmant/models"
)

func main() {

	db := make(map[string]*models.TeamMember)

	alice := &models.TeamMember{Name: "Alice"}
	boob := &models.TeamMember{Name: "Boob"}

	db["Alice"] = alice
	db["Boob"] = boob

	alice.AddTask("taskManagmant")
	alice.AddTask("Change password")

	alice.UpdateTaskStatus("taskManagmant", "done")

	fmt.Println("Tasks *********  ")

	DisplayAllTasks(db["Alice"])

	DisplayAllMembers(db)

}

func DisplayAllTasks(member *models.TeamMember) {

	if len(member.Tasks) > 0 {

		fmt.Printf("%s tasks:\n", member.Name)

		for i, task := range member.Tasks {

			fmt.Printf("task %d : %+v\n", i+1, task)
		}
	} else {
		fmt.Printf("memeber %s didn't have any task yet\n", member.Name)

	}
}

func DisplayAllMembers(db map[string]*models.TeamMember) {

	for _, member := range db {
		fmt.Printf("Member: %s\n", member.Name)
		DisplayAllTasks(member)

	}
}
