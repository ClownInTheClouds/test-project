package main

import (
	"fmt"

	"test-project/internal/repository"
)

func main() {
	var taskRepository = repository.CsvTaskRepository{Path: []string{"data", "tasks.csv"}}

	tasks, err := taskRepository.ReadAll()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, task := range tasks {
			fmt.Println(task)
		}
	}
}
