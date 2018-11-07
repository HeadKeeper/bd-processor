package main

import (
	"./stackoverflow"
	"log"
	"sync"
)

const _WAIT_GROUP_SIZE = 1

func main() {
	log.Println("Starting data fetching...")
	var waitGroup sync.WaitGroup
	waitGroup.Add(_WAIT_GROUP_SIZE)
	go stackoverflow.ParseData("c++", &waitGroup)
	//go stackoverflow.StartFetching(&waitGroup)
	//go github.StartFetching(&waitGroup)

	waitGroup.Wait()
}
