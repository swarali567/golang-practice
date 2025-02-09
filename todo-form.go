package main

import (
	"fmt"
	"net/http"
	"strings"
)

var taskitems = []string{
	"Get fruits", 
	"Pay lightbill", 
	"Watch Golang course", 
	"Treat with cheesecake",
}

func main() {
	// Print a welcome message
	fmt.Println("***Welcome to the To-Do List App!***")

	// HTTP handler to display the To-Do list
	http.HandleFunc("/", todoListHandler)

	// HTTP handler to add a task
	http.HandleFunc("/add", addTaskHandler)

	// Start the server
	http.ListenAndServe(":8080", nil)
}

// Handler function to display the To-Do list
func todoListHandler(w http.ResponseWriter, r *http.Request) {
	var builder strings.Builder
	builder.WriteString("<h1>Your To-Do List:</h1>")
	builder.WriteString("<ul>")

	// Add each task to the response
	for index, task := range taskitems {
		builder.WriteString(fmt.Sprintf("<li>%d. %s</li>", index+1, task))
	}

	builder.WriteString("</ul>")
	builder.WriteString("<form method='POST' action='/add'>")
	builder.WriteString("<input type='text' name='task' placeholder='New Task' required/>")
	builder.WriteString("<button type='submit'>Add Task</button>")
	builder.WriteString("</form>")

	// Write the response to the browser
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(builder.String()))
}

// Handler function to add a task
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Get the task from the form input
		r.ParseForm()
		task := r.FormValue("task")

		// Add the task to the list
		taskitems = append(taskitems, task)

		// Redirect back to the todo list
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
