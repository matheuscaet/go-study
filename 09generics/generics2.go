package generics

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

type OldTask struct {
	Name      string
	Completed bool
}

func Generics2() {
	fmt.Println("-------------------------------- 09 generics2 starts --------------------------------")

	tasks, err := loadJson[Task]("./files/tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	oldTasks, err := loadJson[OldTask]("./files/oldTasks.json")
	if err != nil {
		fmt.Println("Error loading old tasks:", err)
	}

	fmt.Println(tasks)
	fmt.Println(oldTasks)
	fmt.Println("-------------------------------- 09 generics2 ends --------------------------------")
}

func loadJson[T Task | OldTask](filename string) ([]T, error) {
	var tasks []T

	jsonFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonFile, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
