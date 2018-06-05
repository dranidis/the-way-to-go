package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

// WORKER OMIT
func worker(i int) {
	log.Printf("Starting worker %d\n", i)
	for t := range taskch { // while there are tasks to read from the channel
		t.Work(i)
	}
	done <- true
}

// END OMIT

// Task OMIT
type Task struct {
	msg string
	tms int // time to complete in ms
}

// END OMIT

// Work OMIT
func (t Task) Work(i int) {
	log.Printf("Worker %d: working on task %s for %d ms\n", i, t.msg, t.tms)
	time.Sleep(time.Duration(t.tms) * time.Millisecond)
	log.Printf("Worker %d: finished task %s\n", i, t.msg)
}

func createTasks(noTasks int) []Task {
	tasks := []Task{}

	for i := 0; i < noTasks; i++ {
		tname := fmt.Sprintf("Task %d", i)
		tms := 90 + rand.Intn(20)
		tasks = append(tasks, Task{tname, tms})
	}
	return tasks
}

// CHAN OMIT
var taskch chan Task
var done chan bool

// END OMIT

// MAIN OMIT
func main() {
	log.SetOutput(os.Stdout)
	log.SetOutput(ioutil.Discard)
	start := time.Now()

	taskch = make(chan Task)
	done = make(chan bool)

	noWorkers := 2

	for i := 0; i < noWorkers; i++ {
		go worker(i)
	}
	for _, t := range createTasks(10) {
		taskch <- t
	}
	close(taskch)
	for i := 0; i < noWorkers; i++ {
		<-done
	}

	fmt.Printf("Execution took %s\n", time.Since(start))
}

// END OMIT
