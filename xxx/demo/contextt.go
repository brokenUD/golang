package demotest

import (
	"context"
	"fmt"
	"time"
)

func TimeoutContextCancel() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*800))
	go func() {
		defer cancel()
	}()

	select {
	case <- ctx.Done():
		fmt.Printf("time out")
		return
	}
}
