package event

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateEventCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	Generate(ctx)
	go func() {
		time.Sleep(1 * time.Millisecond)
		if true { // Set false to cause test to fail
			cancel()
		}
	}()
	select {
	case <-ctx.Done():
	case <-time.After(10 * time.Millisecond):
		assert.FailNow(t, "Cancelled failed")
	}
}
