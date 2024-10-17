package main

import (
	"fmt"
)

func main() {
	shortGolang := "Watch Go crash course"
	fullGolang := "Watch Nana's Golang Full Course"
	rewardDessert := "Reward myself with rewardDessert donut"

	taskItems := []string{shortGolang, fullGolang, rewardDessert}

	fmt.Println("##### Welcome to our Todolist App! #####")

	fmt.Println()
	printTasks(taskItems)

	taskItems = addTask(taskItems, "Go for a run")
	taskItems = addTask(taskItems, "Practice coding in Go")

	fmt.Println()
	fmt.Println("Updated list")
	printTasks(taskItems)
}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
