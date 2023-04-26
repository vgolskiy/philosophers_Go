package main

import (
	"sync"
	"time"
)

func actions(ph *Philosopher, s *Status) {
	currentTime := time.Now().UnixNano()
	ph.LifeLimit = currentTime + s.TimeDie

	for {
		select {
		case <-s.Stop:
			printMessage(s, "dead")
			break
		default:
			takeForks(ph, s)
			eatFood(ph, s)
			placeForks(ph, s)
			printMessage(s, "thinking")
		}
	}
	return
}

func parallelize(s *Status) error {
	var wg sync.WaitGroup

	wg.Add(s.Quantity)
	for i := 0; i < s.Quantity; i++ {
		go func(ph *Philosopher, s *Status) {
			defer wg.Done()
			actions(ph, s)
		}(&s.Philosophers[i], s)
		time.Sleep(time.Microsecond * 42)
	}

	wg.Wait()
	return nil
}
