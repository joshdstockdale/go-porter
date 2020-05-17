package main

import "go-porter/porter"

func main(){
	m := porter.GetManager(10)
	m.AssembleWorkers()
}
