/*
Implement the dining philosopher’s problem with the following constraints/modifications.

- There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
- Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
- The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
- In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
- The host allows no more than 2 philosophers to eat concurrently.
- Each philosopher is numbered, 1 through 5.
- When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
- When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
Submission: Upload your source code for the program.
*/

package main

import (
	"fmt"
	"reflect"
	"sync"
)

var PHILOSOPHERS_NUM = 5
var INITIAL_HUNGER = 3
var CONCURRENT_EATERS_LIMIT = 2
var TOTAL_MEALS_TO_BE_SERVED = PHILOSOPHERS_NUM * INITIAL_HUNGER

var START_PARTY = "The host has started the dinner party."
var END_PARTY = "The dinner party is over. Everyone go home!"
var STARTED_EATING_MSG = "Starting to eat '%d'.\n"
var EATING_MSG = "nomnomnom... '%d'.\n"
var FINISHED_EATING_MSG = "Finishing eating '%d'.\n"

var LEFT = "left"
var RIGHT = "right"

type Chopstick struct {
	id     int
	isFree *sync.Mutex
}

type Philosopher struct {
	chopsticks map[string]Chopstick
	id         int
	hunger     int
}

func (philo Philosopher) IsHungry() bool {
	return philo.hunger > 0
}

func (philo Philosopher) CanEat() bool {
	return philo.IsHungry() && MutexUnlocked(philo.chopsticks[LEFT].isFree) && MutexUnlocked(philo.chopsticks[RIGHT].isFree)
}

func (philo *Philosopher) Eat(wg *sync.WaitGroup, eatingChannel chan<- int) {
	defer wg.Done()
	fmt.Printf(STARTED_EATING_MSG, philo.id)
	philo.chopsticks[LEFT].isFree.Lock()
	philo.chopsticks[RIGHT].isFree.Lock()
	fmt.Printf(EATING_MSG, philo.id)
	philo.hunger -= 1
	philo.chopsticks[LEFT].isFree.Unlock()
	philo.chopsticks[RIGHT].isFree.Unlock()
	eatingChannel <- philo.id
}

type Host struct {
	philosophers []*Philosopher
}

func (host *Host) StartParty() {
	var wg sync.WaitGroup

	fmt.Println(START_PARTY)

	host.Feed(&wg)

	wg.Wait()

	fmt.Println(END_PARTY)
}

func (host Host) HasHungryGuest() bool {
	for _, philo := range host.philosophers {
		if philo.CanEat() {
			return true
		}
	}
	return false
}

func (host Host) Feed(wg *sync.WaitGroup) {
	for _, philo := range host.philosophers {
		if philo.CanEat() {
			var eatingChannel = make(chan int, CONCURRENT_EATERS_LIMIT)
			wg.Add(1)

			go philo.Eat(wg, eatingChannel)

			id := <-eatingChannel
			fmt.Printf(FINISHED_EATING_MSG, id)
			host.Feed(wg)
		}
	}
}

const MUTEX_LOCKED = 0

func MutexUnlocked(m *sync.Mutex) bool {
	state := reflect.ValueOf(m).Elem().FieldByName("state")
	return state.Int()&MUTEX_LOCKED == MUTEX_LOCKED
}

func makeChopstick(id int) Chopstick {
	return Chopstick{
		id:     1,
		isFree: &sync.Mutex{},
	}
}

func makePhilosopher(id int, leftChopstick Chopstick, rightChopstick Chopstick) Philosopher {
	return Philosopher{
		chopsticks: map[string]Chopstick{
			LEFT:  leftChopstick,
			RIGHT: rightChopstick,
		},
		id:     id,
		hunger: INITIAL_HUNGER,
	}
}

func makeHost(philosphers []*Philosopher) Host {
	return Host{
		philosophers: philosphers,
	}
}

func main() {
	c1 := makeChopstick(1)
	c2 := makeChopstick(2)
	c3 := makeChopstick(3)
	c4 := makeChopstick(4)
	c5 := makeChopstick(5)
	p1 := makePhilosopher(1, c1, c2)
	p2 := makePhilosopher(2, c2, c3)
	p3 := makePhilosopher(3, c3, c4)
	p4 := makePhilosopher(4, c4, c5)
	p5 := makePhilosopher(5, c5, c1)

	philosophers := []*Philosopher{}

	philosophers = append(philosophers, &p1)
	philosophers = append(philosophers, &p2)
	philosophers = append(philosophers, &p3)
	philosophers = append(philosophers, &p4)
	philosophers = append(philosophers, &p5)

	host := makeHost(philosophers)

	host.StartParty()
}
