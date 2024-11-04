package models

import "fmt"

type TeamMember struct {
	Name  string
	Tasks []*Task
}

func (member *TeamMember) AddTask(title string) {

	task := &Task{Title: title, Status: "Not Started"}
	member.Tasks = append(member.Tasks, task)

}

func (member *TeamMember) UpdateTaskStatus(taskTitle string, newStatus string) bool {

	// si aun no hay tareas en el slice devuelve false
	if len(member.Tasks) == 0 {
		return false
	}

	for _, task := range member.Tasks {

		//Si la tarea existe devuelve true
		if task.Title == taskTitle {

			task.Status = newStatus
			fmt.Printf("The status of the task %s has change", taskTitle)
			return true
		}
	}

	//En caso de no encontrar la tarea devuelve false
	fmt.Printf("The task with a title %s was not fund!", taskTitle)
	return false

}
