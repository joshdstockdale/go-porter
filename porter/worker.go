package porter

import (
	"fmt"
	"github.com/buger/jsonparser"
)

type worker struct{
}

func (w *worker) DoTheThing(){
	t := Task{
		Id: 1,
		Thing: []byte(`{"post": {"title": "New Title"}}`),
	}
	t.convertThing()
}

func (tk *Task) convertThing(){
	c, _, _,err := jsonparser.Get(tk.Thing, "post")
	if err != nil{
		fmt.Println("Error:", err)
	}
	fmt.Println("New Title is ", c)
}

func (w *worker) ReportBack(){
	fmt.Println("Report Back")
}