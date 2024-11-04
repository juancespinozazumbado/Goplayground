package models

import "fmt"

func DisplayAllTasks(member *TeamMember) {

	if len(member.Tasks) > 0 {

		fmt.Printf("%s tasks:\n", member.Name)

		fmt.Printf("n | Title | Status | \n")

		for i, task := range member.Tasks {

			fmt.Printf(" %d | %s | %s\n", i+1, task.Title, task.Status)
		}
	} else {
		fmt.Printf("memeber %s didn't have any task yet\n", member.Name)

	}
}

func DisplayAllMembers(db map[string]*TeamMember) {

	for _, member := range db {
		fmt.Printf("Name: %s | TotalTasks: %d\n", member.Name, len(member.Tasks))

	}
}
