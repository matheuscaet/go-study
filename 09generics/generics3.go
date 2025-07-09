package generics

import "fmt"

type TaskNew struct {
	Name      string
	Completed bool
}

type TaskOld struct {
	Description string
	Done        bool
}

type TaskType[T TaskNew | TaskOld] struct {
	ID          int
	Definitions T
}

func Generics3() {
	fmt.Println("-------------------------------- 09 generics3 starts --------------------------------")

	taskNew := TaskType[TaskNew]{
		ID: 1,
		Definitions: TaskNew{
			Name:      "Buy groceries",
			Completed: false,
		},
	}

	taskOld := TaskType[TaskOld]{
		ID: 2,
		Definitions: TaskOld{
			Description: "Buy a new phone",
			Done:        false,
		},
	}

	fmt.Println(taskNew)
	fmt.Println(taskOld)

	fmt.Println("-------------------------------- 09 generics3 ends --------------------------------")
}
