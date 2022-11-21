package task

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Task struct {
	Date time.Time
	Desc string
}

func (t *Task) Add(d string) (Task, error) {
	task := Task{
		Date: time.Now(),
		Desc: d,
	}
	return task, nil
}

func (t *Task) Store() error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("task.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) Load() ([]Task, error) {
	var tasks []Task
	data, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		return []Task{}, nil
	}
	err = json.Unmarshal(data, tasks)
	if err != nil {
		return []Task{}, nil
	}
	return tasks, nil
}
