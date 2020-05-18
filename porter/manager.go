package porter

import (
	"time"
)

func WakePorters(maxWorkers int, tasks []Task)  {
	mw := len(tasks)
	if mw > maxWorkers {
		mw = maxWorkers
	}

	//divide tasks by mw, creating 2 groups. One is limited, Two is not
	// g := append(tasks, tasks[:mw-1]); g = append(tasks, tasks[mw:]
	//open the threads, go func() up to the limit of mw
	//loop through threads and pass func in
	//wait for a thread to finish, pass another func in freed up thread

	for _,t := range tasks{
		//Added in for debugging
		time.Sleep(1000 * time.Millisecond)

		w := Worker{
			Id: time.Now(),
			Task: t,
		}
		w.DoTheThing()

	}

}
