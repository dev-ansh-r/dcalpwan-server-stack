package component_test

import (
	"context"
	"sync"
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/task"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestTasks(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	c, err := New(test.GetLogger(t), &Config{})
	a.So(err, should.BeNil)

	// Register a one-off task.
	oneOffWg := sync.WaitGroup{}
	oneOffWg.Add(1)
	c.RegisterTask(&task.Config{
		Context: ctx,
		ID:      "one_off",
		Func: func(_ context.Context) error {
			oneOffWg.Done()
			return nil
		},
		Restart: task.RestartNever,
		Backoff: TestTaskBackoffConfig,
	})
	// Register an always restarting task.
	restartingWg := sync.WaitGroup{}
	restartingWg.Add(5)
	i := 0
	c.RegisterTask(&task.Config{
		Context: ctx,
		ID:      "restarts",
		Func: func(_ context.Context) error {
			i++
			if i <= 5 {
				restartingWg.Done()
			}
			return nil
		},
		Restart: task.RestartAlways,
		Backoff: TestTaskBackoffConfig,
	})

	// Register a task that restarts on failure.
	failingWg := sync.WaitGroup{}
	failingWg.Add(1)
	j := 0
	c.RegisterTask(&task.Config{
		Context: ctx,
		ID:      "restarts_on_failure",
		Func: func(_ context.Context) error {
			j++
			if j < 5 {
				return errors.New("failed")
			}
			failingWg.Done()
			return nil
		},
		Restart: task.RestartOnFailure,
		Backoff: TestTaskBackoffConfig,
	})

	// Wait for all invocations.
	test.Must[any](nil, c.Start())
	defer c.Close()
	oneOffWg.Wait()
	restartingWg.Wait()
	failingWg.Wait()
}
