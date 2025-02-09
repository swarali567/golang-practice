package main

import (
	"fmt"
	"net/http"
)

var taskitems = []string{"Get fruits", "Pay lightbill", "Watch Golang course", "Treat with cheesecake"}

func main() {
	http.HandleFunc("/show-task", helloUser)
	http.HandleFunc("/", showtasks)

	http.ListenAndServe(":8080", nil)
	printtask(taskitems)
	fmt.Println()
	taskitems = addtask(taskitems, "Go fo run")
	fmt.Println("Updated List")
	printtask(taskitems)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "***Welcome to the To-Do List App!***"
	fmt.Fprintln(writer, greeting)
}

func showtasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskitems {
		fmt.Fprintln(writer, task)
	}
}

func printtask(taskitems []string) {
	for index, task := range taskitems {
		fmt.Println(index+1, ".", task)
	}
}

func addtask(taskitems []string, newtask string) []string {
	var updatetask = append(taskitems, newtask)
	return updatetask
}
