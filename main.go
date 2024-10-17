package main

import (
	"fmt"
	"log"
	"net/http"
)

var shortGolang = "Watch Go crash course"
var fullGolang = "Watch Nana's Golang Full Course"
var rewardDessert = "Reward myself with rewardDessert donut"

var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func showTasks(writer http.ResponseWriter, _ *http.Request) {
	for _, task := range taskItems {
		_, err := fmt.Fprintln(writer, task)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func helloUser(writer http.ResponseWriter, _ *http.Request) {
	greeting := "Hello, User! Welcome to our Todolist App!"
	_, err := fmt.Fprintf(writer, greeting)
	if err != nil {
		log.Fatal(err)
	}
}
