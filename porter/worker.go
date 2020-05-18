package porter

import (
	"fmt"
	"github.com/buger/jsonparser"
	"time"
)

type Worker struct{
	Id time.Time
	Task Task
}

func (w *Worker) DoTheThing(){
	c,err := jsonparser.GetString(w.Task.Package, "title")
	if err != nil{
		fmt.Println("Error:", err)
	}

	fmt.Printf("Task for %d is %v \n", w.Task.Id, c)
}


func (w *Worker) ReportBack(){
	fmt.Println("Report Back")
}