package utilities

import "time"


type Game struct {
	NumRounds  int
	Level      int
	Win        int
	Loss       int
	Result     int
	Answer     int
	Round      []int
	Op         string
	Operations []string
	RoundStart time.Time
	RoundTime  time.Duration
	TotalTime  time.Duration
	Err        error
}


