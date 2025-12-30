package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
var scanner = bufio.NewScanner(os.Stdin)

func main() {

	for true {

		showMenu()

		var choice int
		fmt.Println("Enter your choice (1 - 9): ")
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		checkError(err)

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
	tasks := loadTasks()
	var ts Task

	fmt.Printf("Input task description: ")
	scanner.Scan()
	ts.Desc = scanner.Text()

	if len(tasks) == 0 {
		ts.ID = 1;
	} else {
		ts.ID = tasks[len(tasks) - 1].ID + 1
	}

	ts.Status = 0
	ts.Create = time.Now().Format(time.RFC3339)
	ts.Update = time.Now().Format(time.RFC3339)

	tasks = append(tasks, ts)
	saveFile(tasks)
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
