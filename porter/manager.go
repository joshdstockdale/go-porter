package porter

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
	maxWorkers int
}

var singleInstance *single

func GetManager(mw int) *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating Porter Manager")
				singleInstance = &single{mw}
			})
	} else {
		fmt.Println("You can only serve one manager.")
	}
	return singleInstance
}

func (s *single) AssembleWorkers() []*worker {
	wks := make([]*worker, s.maxWorkers-1, s.maxWorkers-1)
	return wks
}
