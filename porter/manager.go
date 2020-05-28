package porter

import (
	"fmt"
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

	// divide tasks by mw, creating 2 groups. One is limited to mw, Two is not
	g := make([][]Task, 2, 2)
	g[0] = tasks[:mw]
	g[1] = tasks[mw:]
	length := len(g[1])-1
	// open the threads, go func() up to the limit of mw

	// loop through threads and pass func in
	// wait for a thread to finish, pass another func in freed up thread

	// Add WaitGroup
	wg.Add(1)

	//Do group 1
	for i,t := range g[0]{
		// Added in for debugging
		w := Worker{
			Id: time.Now(),
			Task: t,
		}
		go w.DoTheThing(&wg)
		if(i < length ){
			wg.Add(1)
		}
	}
	fmt.Println("Group 1 done")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()


	//Do group 2
	//for i,t := range g[1]{
	//	//Added in for debugging
	//	w := Worker{
	//		Id: time.Now(),
	//		Task: t,
	//	}
	//	go w.DoTheThing()
	//	wg.Done()
	//	fmt.Printf("Group 2 #%v done",i)
	//	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	//	if(i < len(g[0])-1 ){
	//		wg.Add(1)
	//	}
	//}

	fmt.Println("ALL done")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func oneByOneTasks(wg *sync.WaitGroup, t Task, i int, length int){
	// do group 2, one at a time
	wg.Wait()
	fmt.Println("oneByOneTasks")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	w := Worker{
		Id: time.Now(),
		Task: t,
	}
	w.DoTheThing(wg)
	if i < length {
		wg.Add(1)
	}

}