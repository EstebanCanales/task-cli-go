package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Status string

const (
	Done       Status = "Done"
	NoDone     Status = "No done"
	InProgress Status = "In Progress"
)

type Task struct {
	ID   int
	Name string
	Done Status
}

// var tasks []Task // Removed global tasks slice

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		status TEXT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

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

	stmt, err := db.Prepare("INSERT INTO tasks(name, status) VALUES(?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, NoDone)
	if err != nil {
		log.Println("Error executing statement:", err)
		return
	}

	fmt.Println("Task added successfully!")
	ClearScreen()
}

func deleteTask() {
	showTasks() // Show tasks so user can see IDs

	var id int
	fmt.Print("Enter the task ID to delete: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		log.Println("Invalid input for ID:", err)
		return
	}

	stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		log.Println("Error executing statement:", err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Printf("No task found with ID %d.\n", id)
	} else {
		fmt.Printf("Task with ID %d deleted successfully!\n", id)
	}
	ClearScreen()
}

func showTasks() {
	rows, err := db.Query("SELECT id, name, status FROM tasks")
	if err != nil {
		log.Println("Error querying tasks:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\nTasks:")
	hasTasks := false
	for rows.Next() {
		hasTasks = true
		var task Task
		err := rows.Scan(&task.ID, &task.Name, &task.Done)
		if err != nil {
			log.Println("Error scanning task row:", err)
			continue
		}
		fmt.Printf("%d. %s - %s\n", task.ID, task.Name, task.Done)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating tasks rows:", err)
		return
	}

	if !hasTasks {
		fmt.Println("No tasks in the list.")
	}
}

func markInProgress() {
	showTasks() // Show tasks so user can see IDs

	var id int
	fmt.Print("Enter the task ID to mark as in progress: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		log.Println("Invalid input for ID:", err)
		return
	}

	stmt, err := db.Prepare("UPDATE tasks SET status = ? WHERE id = ?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(InProgress, id)
	if err != nil {
		log.Println("Error executing statement:", err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Printf("No task found with ID %d.\n", id)
	} else {
		fmt.Printf("Task with ID %d marked as in progress successfully!\n", id)
	}
	ClearScreen()
}

func markDone() {
	showTasks() // Show tasks so user can see IDs

	var id int
	fmt.Print("Enter the task ID to mark as done: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		log.Println("Invalid input for ID:", err)
		return
	}

	stmt, err := db.Prepare("UPDATE tasks SET status = ? WHERE id = ?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(Done, id)
	if err != nil {
		log.Println("Error executing statement:", err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Printf("No task found with ID %d.\n", id)
	} else {
		fmt.Printf("Task with ID %d marked as done successfully!\n", id)
	}
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
