package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Status string

const (
	Done       Status = "Done"
	NoDone     Status = "No done"
	InProgress Status = "In Progress"
)

type Task struct {
	Name string
	Done Status
}

var tasks []Task

func main() {
	fmt.Println("Task CLI")
	for {
		fmt.Println("\nSelect Option")
		fmt.Println("1. Add Task")
		fmt.Println("2. Show Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Mark in progress")
		fmt.Println("5. Mark done")
		fmt.Println("6. Quit")

		var option int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			addTask()
		case 2:
			showTasks()

			fmt.Println("\nDo you want to exit the list? (yes/no)")

			if yesNo() {
				ClearScreen()
				break
			}

		case 3:
			deleteTask()
		case 4:
			markInProgress()
		case 5:
			markDone()
		case 6:
			ClearScreen()
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid Option. Please try again.")
		}
	}
}

func addTask() {
	var name string
	fmt.Print("Enter task name: ")
	fmt.Scanln(&name)
	tasks = append(tasks, Task{Name: name, Done: NoDone})
	fmt.Println("Task added successfully!")
	ClearScreen()
}

func deleteTask() {
	if len(tasks) == 0 {
		fmt.Println("No tasks in the list to delete.")
		return
	}

	showTasks()

	var name string
	fmt.Print("Enter the task name to delete: ")
	fmt.Scanln(&name)

	for i, task := range tasks {
		if task.Name == name {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Task '%s' deleted successfully!\n", name)
			return
		}
	}

	fmt.Printf("Task '%s' not found.\n", name)
	ClearScreen()
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}

	fmt.Println("\nTasks:")
	for i := range tasks {
		fmt.Printf("%d. %s - %s\n", i+1, tasks[i].Name, tasks[i].Done)
	}
}

func markInProgress() {
	if len(tasks) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}

	showTasks()

	var name string
	fmt.Print("Enter the task name to mark as in progress: ")
	fmt.Scanln(&name)

	for i, task := range tasks {
		if task.Name == name {
			tasks[i].Done = InProgress
			fmt.Printf("Task '%s' marked as in progress successfully!\n", name)
			return
		}
	}

	fmt.Printf("Task '%s' not found.\n", name)
	ClearScreen()
}

func markDone() {
	if len(tasks) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}

	showTasks()

	var name string
	fmt.Print("Enter the task name to mark as done: ")
	fmt.Scanln(&name)

	for i, task := range tasks {
		if task.Name == name {
			tasks[i].Done = Done
			fmt.Printf("Task '%s' marked as done successfully!\n", name)
			return
		}
	}

	fmt.Printf("Task '%s' not found.\n", name)
	ClearScreen()
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func yesNo() bool {
	for {
		var answer string
		fmt.Scanln(&answer)
		return answer == "yes" || answer == "y"
	}
}
