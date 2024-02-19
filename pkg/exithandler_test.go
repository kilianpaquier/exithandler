package exithandler_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	exithandler "github.com/kilianpaquier/exithandler/pkg"
)

func TestExitHandler(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		// Arrange
		ctx, cancel := context.WithTimeout(ctx, 5*time.Millisecond)
		t.Cleanup(cancel)
		check := atomic.Bool{}
		check.Store(false)

		// Act
		go exithandler.Handle(ctx, func(context.Context) {
			check.Store(true)
		})
		<-ctx.Done()

		// Assert
		assert.True(t, check.Load())
	})

	t.Run("success_delayed", func(t *testing.T) {
		// Arrange
		ctx, cancel := context.WithTimeout(ctx, 5*time.Millisecond)
		t.Cleanup(cancel)
		check := atomic.Bool{}
		check.Store(false)

		// Act
		go exithandler.HandleFunc(ctx, func(context.Context) {
			check.Store(true)
		})()
		<-ctx.Done()

		// Assert
		assert.True(t, check.Load())
	})
}
