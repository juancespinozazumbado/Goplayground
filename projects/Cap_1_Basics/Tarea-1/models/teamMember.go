package models

type TeamMember struct {
	Name  string
	Tasks []*Task
}

func (member *TeamMember) AddTask(title string) {

	task := &Task{Title: title, Status: "Not Started"}
	member.Tasks = append(member.Tasks, task)

}

func (member *TeamMember) UpdateTaskStatus(taskTitle string, newStatus string) bool {

	if len(member.Tasks) == 0 {
		return false
	}

	for _, task := range member.Tasks {

		if task.Title == taskTitle {

			task.Status = newStatus
			return true
		}
	}

	return false

}
