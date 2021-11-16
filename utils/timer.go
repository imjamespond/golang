package utils

import "time"

func StartTicker(seconds int64, do func()) *time.Ticker {
	ticker := time.NewTicker(time.Second * time.Duration(seconds))
	go func(t *time.Ticker) {
		do()
		for range t.C {
			do()
		}
	}(ticker)
	return ticker
}
