package main

type Philosopher struct {
	PhilosopherID int
	CheckerID     int
	Number        int
	Eating        bool
	EatQty        int
	ForkLeft      int
	ForkRight     int
	LifeLimit     int64
}

func initPhilosophers(s *Status) {
	s.Philosophers = make([]Philosopher, s.Quantity)

	for i := 0; i < s.Quantity; i++ {
		s.Philosophers[i].Number = i
		s.Philosophers[i].ForkLeft = i
		s.Philosophers[i].ForkRight = (i + 1) % s.Quantity
	}
	return
}
