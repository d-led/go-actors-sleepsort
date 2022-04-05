package interfaces

import (
	"molizen-sleepsort/actors"

	"github.com/sanposhiho/molizen/context"
)

type Waiter interface {
	WaitFor(ctx context.Context, delaySec int)
	SetSelf(ctx context.Context, self *actors.WaiterActor)
	Print(ctx context.Context)
}
