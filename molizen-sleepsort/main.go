//go:generate go install github.com/sanposhiho/molizen/cmd/molizen@latest
//go:generate molizen -source interfaces/waiter.go -destination actors/waiter.go -package actors

package main

import (
	"fmt"
	"math/rand"
	"molizen-sleepsort/actors"
	"os"
	"path"
	"sync"
	"time"

	"github.com/sanposhiho/molizen/context"

	"github.com/sanposhiho/molizen/actor"

	"github.com/sanposhiho/molizen/node"
)

func main() {
	whereami, _ := os.Getwd()
	fmt.Printf("--=== %v ===--\n", path.Base(whereami))
	rand.Seed(time.Now().UnixMicro())

	node := node.NewNode()
	ctx := node.NewContext()

	count := 10

	numbers := []int{}
	workers := []*actors.WaiterActor{}

	// prepare work and workers
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		numbers = append(numbers, rand.Intn(15))
		actorFuture := actors.New(ctx, &waiter{
			wg: &wg,
		}, actor.Option{})
		actor := actorFuture.Get(ctx).Actor
		future := actor.SetSelf(ctx, &actor)
		future.Get(ctx)
		workers = append(workers, &actor)
	}
	wg.Add(count)

	// start
	for i := 0; i < count; i++ {
		workers[i].WaitFor(ctx, numbers[i])
	}

	wg.Wait()
}

type waiter struct {
	wg       *sync.WaitGroup
	delaySec int
	self     *actors.WaiterActor
}

func (w *waiter) WaitFor(ctx context.Context, delaySec int) {
	w.delaySec = delaySec
	time.Sleep(time.Duration(delaySec) * time.Second)
	w.self.Print(ctx)
}

func (w *waiter) SetSelf(ctx context.Context, self *actors.WaiterActor) {
	w.self = self
}

func (w *waiter) Print(ctx context.Context) {
	fmt.Println(w.delaySec)
	w.wg.Done()
}
