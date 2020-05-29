package porter

import (
	"fmt"
	"github.com/buger/jsonparser"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func WakePorters(maxWorkers int, tasks []Task)  {
	mw := len(tasks)
	if mw > maxWorkers {
		mw = maxWorkers
	}

	wg.Add(len(tasks))
	//Do group 1
	for i,t := range tasks{

		// Added in for debugging
		w := Worker{
			Id: time.Now(),
			Task: t,
		}
		go func() {
			c,err := jsonparser.GetString(w.Task.Package, "title")
			if err != nil{
				fmt.Println("Error:", err)
			}

			time.Sleep(2 * time.Second)
			fmt.Printf("Task for %d is %v \n", w.Task.Id, c)
			wg.Done()
		}()
		if(i < len(tasks)-1 ){
			wg.Add(1)
		}
		fmt.Println(i)
		if(i >= mw-1){
			wg.Wait()
		}
	}
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("ALL done")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}
