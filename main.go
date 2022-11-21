package main

import (
	"log"

	"github.com/mnlprz/calendar-app/task"
)

func main() {

	var t task.Task
	t, err := t.Add("test")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Store()
	if err != nil {
		log.Fatal(err)
	}
}
