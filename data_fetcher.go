package main

import (
	"log"
	"errors"
)

const _STACKOVERFLOW_DATA_FETCHER = "stackoverflow"
const _GITHUB_DATA_FETCHER = "github"
const STACKOVERFLOW_URL_PATTERN = ""
const GITHUB_URL_PATTERN  = ""

// Every data fetcher works with his own file
// Data fetcher's lifecycle - init, start

type DataFetcher struct {
	target string
	urlPattern string
}

type fetchFunc func()
type urlPreparingFunc func()

/*
	Factory for getting data fetcher's
*/
func GetDataFetcher(target string) (*DataFetcher, error) {
	dataFetcher := new(DataFetcher)
	switch target {
		case _STACKOVERFLOW_DATA_FETCHER:
			dataFetcher.target = "github"
			dataFetcher.urlPattern =
		case _GITHUB_DATA_FETCHER:
			dataFetcher.target = "stackoverflow"
			dataFetcher.urlPattern =
		default:
			return nil, errors.New("Invalid fetcher type.")
	}
	return dataFetcher, nil
}

func (fetcher *DataFetcher) Initialize() {
	// init steps
	log.Printf("Initialized new data fetcher with target %s.\n", fetcher.target)
}

func (fetcher *DataFetcher) Start(fetch fetchFunc) {

}