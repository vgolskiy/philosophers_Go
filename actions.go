package main

import (
	"fmt"
	"time"
)

func eatFood(ph *Philosopher, s *Status) {
	ph.Eating = true
	currentTime := time.Now().UnixNano()
	ph.LifeLimit = currentTime + s.TimeDie
	select {
	case <-s.Stop:
		return
	default:
		printMessage(ph, s, "eating")
	}
	time.Sleep(time.Microsecond * time.Duration(s.TimeEat))
	ph.EatQty++
	ph.Eating = false
	return
}

func takeForks(ph *Philosopher, s *Status) {
	s.MutexForks[ph.ForkLeft].Lock()
	select {
	case <-s.Stop:
		return
	default:
		printMessage(ph, s, "took left fork")
	}
	s.MutexForks[ph.ForkRight].Lock()
	select {
	case <-s.Stop:
		return
	default:
		printMessage(ph, s, "took right fork")
	}
	return
}

func placeForks(ph *Philosopher, s *Status) {
	select {
	case <-s.Stop:
		return
	default:
		printMessage(ph, s, "sleeping")
	}
	s.MutexForks[ph.ForkLeft].Unlock()
	s.MutexForks[ph.ForkRight].Unlock()
	time.Sleep(time.Microsecond * time.Duration(s.TimeSleep))
	return
}

func printMessage(ph *Philosopher, s *Status, str string) {
	s.MutexPrint.Lock()
	currentTime := time.Now().UnixNano()
	fmt.Println(fmt.Sprintf("%d: %d %s", currentTime-s.TimeStart, ph.Number+1, str))
	s.MutexPrint.Unlock()
	return
}
