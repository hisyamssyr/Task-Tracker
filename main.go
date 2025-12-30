package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	ID     int		`json:"id"`
	Desc   string	`json:"desc"`
	Status int		`json:"status"`
	Create string	`json:"created_at"`
	Update string	`json:"updated_at"`
}

const filename = "tasks.json"

func main() {

	for true {

		showMenu()

		var choice int
		fmt.Println("Enter your choice (1 - 9): ")
		fmt.Scan(&choice)

		process(choice)
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func showMenu() {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("              TASK TRACKER")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Println("Menu:")
	fmt.Println("1. Add new task")
	fmt.Println("2. Edit existing task")
	fmt.Println("3. Delete existing task")
	fmt.Println("4. Mark done")
	fmt.Println("5. Show all tasks")
	fmt.Println("6. Show all done tasks")
	fmt.Println("7. Show all on progress tasks")
	fmt.Println("8. Show all not done tasks")
	fmt.Println("9. Close program")
}

func process(menu int) {
	switch menu {
	case 1: addTask()
	case 2: editTask()
	case 3: deleteTask()
	case 4: doneTask()
	case 5: showAll()
	case 6: showDone()
	case 7: showOnProgress()
	case 8: showNotDone()
	case 9: close()
	}
}

func loadTasks()([]Task) {
	data, err := os.ReadFile(filename)
	checkError(err)

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	checkError(err)

	return tasks
}

func saveFile(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	checkError(err)

	err = os.WriteFile(filename, data, 0644)
	checkError(err)
}

func addTask() {

}

func editTask() {

}

func deleteTask() {

}

func doneTask() {

}

func showAll() {

}

func showDone() {

}

func showOnProgress() {

}

func showNotDone() {

}

func close() {
	println("Program closed.")
	os.Exit(0)
}

func loadFile() {

}
