package porter

import (
	"fmt"
	"github.com/buger/jsonparser"
	"sync"
	"time"
)

type Worker struct{
	Id time.Time
	Task Task
}

func (w *Worker) DoTheThing(wg *sync.WaitGroup){
	c,err := jsonparser.GetString(w.Task.Package, "title")
	if err != nil{
		fmt.Println("Error:", err)
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("Task for %d is %v \n", w.Task.Id, c)
	wg.Done()
	//Call back, do the next task
}


func (w *Worker) ReportBack(){
	fmt.Println("Report Back")
}