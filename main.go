package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Todo struct
type Todo struct {
	ID        int
	Task      string
	Completed bool
	CreatedAt time.Time
}

var todos []Todo
var nextID = 1

func main() {
	fmt.Println("=== Go To-Do List ===\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		showMenu()
		fmt.Print("Choose option: ")
		
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addTodo(reader)
		case "2":
			listTodos()
		case "3":
			markDone(reader)
		case "4":
			deleteTodo(reader)
		case "5":
			fmt.Println("Goodbye! 👋")
			return
		default:
			fmt.Println("Invalid option! Try again.")
		}
	}
}

func showMenu() {
	fmt.Println("\n1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Mark as Done")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
}

func addTodo(reader *bufio.Reader) {
	fmt.Print("Enter task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	if task == "" {
		fmt.Println("Task cannot be empty!")
		return
	}

	todo := Todo{
		ID:        nextID,
		Task:      task,
		Completed: false,
		CreatedAt: time.Now(),
	}

	todos = append(todos, todo)
	nextID++
	fmt.Println("✅ Task added successfully!")
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No tasks yet! Add some.")
		return
	}

	fmt.Println("\n=== Your Tasks ===")
	for _, todo := range todos {
		status := "❌"
		if todo.Completed {
			status = "✅"
		}
		fmt.Printf("%d. %s %s\n", todo.ID, status, todo.Task)
	}
}

func markDone(reader *bufio.Reader) {
	listTodos()
	if len(todos) == 0 {
		return
	}

	fmt.Print("Enter task ID to mark done: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		fmt.Println("Invalid ID!")
		return
	}

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			fmt.Println("✅ Task marked as done!")
			return
		}
	}
	fmt.Println("Task not found!")
}

func deleteTodo(reader *bufio.Reader) {
	listTodos()
	if len(todos) == 0 {
		return
	}

	fmt.Print("Enter task ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		fmt.Println("Invalid ID!")
		return
	}

	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("🗑️ Task deleted!")
			return
		}
	}
	fmt.Println("Task not found!")
}