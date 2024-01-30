package utilities

import "time"


type Game struct {
	numRounds  int
	level      int
	win        int
	loss       int
	result     int
	answer     int
	round      []int
	op         string
	operations []string
	roundStart time.Time
	roundTime  time.Duration
	totalTime  time.Duration
	err        error
}


