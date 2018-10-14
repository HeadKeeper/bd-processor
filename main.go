package main

import (
	"./stackoverflow"
	"./github"
	"sync"
	"log"
)

const _WAIT_GROUP_SIZE = 2

func main() {
	log.Println("Starting data fetching...")
	var waitGroup sync.WaitGroup
	waitGroup.Add(_WAIT_GROUP_SIZE)
	go stackoverflow.StartFetching(&waitGroup)
	go github.StartFetching(&waitGroup)
	waitGroup.Wait()
}