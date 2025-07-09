package goroutines

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.RWMutex
var data = []string{"one", "two", "three", "four", "five"}
var results = []string{}

func Goroutines() {
	fmt.Println("-------------------------------- 07 goroutines starts --------------------------------")
	t0 := time.Now()

	for _, v := range data {
		wg.Add(1)
		go dbCall(v)
	}

	wg.Wait()

	elapsed := time.Since(t0)
	fmt.Println("Time taken:", elapsed)

	fmt.Println("results", results)

	fmt.Println("-------------------------------- 07 goroutines ends --------------------------------")
}

func dbCall(data string) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	saveData(data)
	logData()
}

func saveData(data string) {
	mutex.Lock()
	results = append(results, data)
	mutex.Unlock()
}

func logData() {
	mutex.RLock()
	fmt.Println("results", results)
	mutex.RUnlock()
}
