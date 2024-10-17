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

	fmt.Println("List of my Todos")
	fmt.Println("Tasks:", taskItems)
}
