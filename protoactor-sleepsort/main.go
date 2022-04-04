package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/scheduler"
)

type sleepFor struct {
	delay time.Duration
}

type print struct {
	delay time.Duration
}

type sleeperActor struct {
	wg *sync.WaitGroup
	s  *scheduler.TimerScheduler
}

func (state *sleeperActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *sleepFor:
		state.s.SendOnce(msg.delay, context.Self(), &print{
			delay: msg.delay,
		})
	case *print:
		fmt.Println(msg.delay.Seconds())
		state.wg.Done()
	}
}

func main() {
	whereami, _ := os.Getwd()
	fmt.Printf("--=== %v ===--\n", path.Base(whereami))

	rand.Seed(time.Now().UnixMicro())

	count := 10
	var wg sync.WaitGroup
	system := actor.NewActorSystem()
	s := scheduler.NewTimerScheduler(system.Root)

	numbers := []int{}
	workers := []*actor.PID{}

	// prepare work and workers
	for i := 0; i < count; i++ {
		numbers = append(numbers, rand.Intn(15))
		props := actor.PropsFromProducer(func() actor.Actor {
			return &sleeperActor{
				wg: &wg,
				s:  s,
			}
		})
		workers = append(workers, system.Root.Spawn(props))
	}
	wg.Add(count)

	for i := 0; i < count; i++ {
		system.Root.Send(workers[i], &sleepFor{
			delay: time.Duration(numbers[i]) * time.Second,
		})
	}

	wg.Wait()
}
