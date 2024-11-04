package main

import (
	"fmt"
	"taskManagmant/models"
)

func main() {

	// Crea la db en memoria
	db := make(map[string]*models.TeamMember)

	//crea team members
	alice := &models.TeamMember{Name: "Alice"}
	boob := &models.TeamMember{Name: "Bob"}
	maria := &models.TeamMember{Name: "Maria"}

	//agrega los members a la db
	db["Alice"] = alice
	db["Boob"] = boob
	db["Maria"] = maria

	showMembers(db)

	//agrega taras a Alice
	alice.AddTask("Write report")
	alice.AddTask("Change password")

	//Agrega tarea a Maria
	db["Maria"].AddTask("Do Homework")

	printTasks(db)

	//Actualiza estado de tarea de Alice
	alice.UpdateTaskStatus("Write report", "done")

	// intenta actualizar el estado de una tarea que no existe
	db["Maria"].UpdateTaskStatus("Rendom Task", "Done")
	//Actualiza estado de tarea Do homework de maria
	db["Maria"].UpdateTaskStatus("Do Homework", "Done")

	printTasks(db)

}

func printTasks(db map[string]*models.TeamMember) {
	fmt.Println("showing all task in the team...: ")

	for _, member := range db {

		models.DisplayAllTasks(member)
	}
}

func showMembers(db map[string]*models.TeamMember) {
	fmt.Println("showing all team members.... ")

	models.DisplayAllMembers(db)
}
