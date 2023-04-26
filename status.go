package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Status struct {
	Quantity     int
	EatMax       int
	Stop         chan bool
	TimeDie      int64
	TimeEat      int
	TimeSleep    int
	TimeStart    int64
	MutexStatus  sync.Mutex
	MutexPrint   sync.Mutex
	MutexForks   []sync.Mutex
	Philosophers []Philosopher
}

func initStatus(args []string, s *Status) error {
	qty, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	//if qty < 2 {
	//	return fmt.Errorf("bring more philosophers")
	//}
	s.Quantity = qty

	s.TimeStart = time.Now().UnixNano()

	die, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}
	if die <= 0 {
		return fmt.Errorf("too short time to death")
	}
	s.TimeDie = int64(die)

	sleep, err := strconv.Atoi(args[4])
	if err != nil {
		return err
	}
	if sleep <= 0 {
		return fmt.Errorf("too short time to sleep")
	}
	s.TimeSleep = sleep

	eat, err := strconv.Atoi(args[3])
	if err != nil {
		return err
	}
	if eat <= 0 {
		return fmt.Errorf("too short time to eat")
	}
	s.TimeEat = eat

	if len(args) == 6 {
		eatMax, err := strconv.Atoi(args[5])
		if err != nil {
			return err
		}
		if eatMax < 2 {
			return fmt.Errorf("wrong maximal eat quantity")
		}
		s.EatMax = eatMax
	}

	s.MutexForks = make([]sync.Mutex, s.Quantity)

	initPhilosophers(s)

	s.Stop = make(chan bool, 1)
	return nil
}
