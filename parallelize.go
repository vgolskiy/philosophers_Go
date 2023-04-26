package main

import (
	"sync"
	"time"
)

func actions(ph *Philosopher, s *Status) {
	for {
		select {
		case <-s.Stop:
			printMessage(ph, s, "dead")
			break
		default:
			takeForks(ph, s)
			eatFood(ph, s)
			placeForks(ph, s)
			printMessage(ph, s, "thinking")
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
			currentTime := time.Now().UnixNano()
			ph.LifeLimit = currentTime + s.TimeDie
			actions(ph, s)
		}(&s.Philosophers[i], s)
		time.Sleep(time.Microsecond * 100)
	}

	wg.Wait()
	return nil
}
