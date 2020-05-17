package main

import "scheduler/porter"

func main(){
	m := porter.GetManager(10)
	m.AssembleWorkers()
}
