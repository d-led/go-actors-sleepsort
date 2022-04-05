package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"sync"
	"time"

	"github.com/Arceliar/phony"
)

type waiter struct {
	phony.Inbox
	delaySec int
	wg       *sync.WaitGroup
}

func (p *waiter) WaitFor(delaySec int) {
	p.Act(p, func() {
		p.delaySec = delaySec
		time.Sleep(time.Duration(delaySec) * time.Second)
		p.Print()
	})
}

func (p *waiter) Print() {
	p.Act(p, func() {
		fmt.Println(p.delaySec)
		p.wg.Done()
	})
}

func main() {
	whereami, _ := os.Getwd()
	fmt.Printf("--=== %v ===--\n", path.Base(whereami))
	rand.Seed(time.Now().UnixMicro())

	count := 10

	numbers := []int{}
	workers := []*waiter{}

	// prepare work and workers
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		numbers = append(numbers, rand.Intn(15))
		workers = append(workers, &waiter{
			wg: &wg,
		})
	}
	wg.Add(count)

	// start
	for i := 0; i < count; i++ {
		workers[i].WaitFor(numbers[i])
	}

	wg.Wait()
}
