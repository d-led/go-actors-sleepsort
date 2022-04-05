package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"sync"
	"time"

	"github.com/ergo-services/ergo"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/ergo-services/ergo/node"
)

type sleeperActor struct {
	gen.Server
	sleeping bool
	wg       *sync.WaitGroup
}

func (s *sleeperActor) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	delay := message.(int)
	// if already scheduled
	if s.sleeping {
		fmt.Println(delay)
		s.wg.Done()
		return gen.ServerStatusStop
	}

	// schedule sleep
	s.sleeping = true
	process.SendAfter(process.Self(), delay, time.Duration(time.Duration(delay)*time.Second))
	return gen.ServerStatusOK
}

func main() {
	// to do: refactor
	whereami, _ := os.Getwd()
	fmt.Printf("--=== %v ===--\n", path.Base(whereami))
	rand.Seed(time.Now().UnixMicro())

	count := 10

	node, _ := ergo.StartNode("node@localhost", "cookie", node.Options{})

	numbers := []int{}
	workers := []etf.Pid{}

	var wg sync.WaitGroup

	// prepare work and workers
	for i := 0; i < count; i++ {
		numbers = append(numbers, rand.Intn(15))
		process, err := node.Spawn(fmt.Sprintf("sleeper-%v", i), gen.ProcessOptions{}, &sleeperActor{
			wg: &wg,
		})
		if err != nil {
			panic(err)
		}
		workers = append(workers, process.Self())
	}
	wg.Add(count)
	fmt.Println(numbers)

	sender, err := node.Spawn("sender", gen.ProcessOptions{}, &sleeperActor{})
	if err != nil {
		panic(err)
	}

	// start
	for i := 0; i < count; i++ {
		err := sender.Send(workers[i], numbers[i])
		if err != nil {
			panic(err)
		}
	}

	wg.Wait()
	node.Stop()
}
