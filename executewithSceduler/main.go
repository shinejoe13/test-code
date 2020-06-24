package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("Task is being performed.")
}

func main() {
	s := gocron.NewScheduler()
	//s.Every(2).Hours().Do(task)
	s.Every(5).Minutes().Do(task)
	<-s.Start()
}
