package concurrent

import (
	"testing"

	"context"

	mock "github.com/mayankzbio/sandbox/sample/concurrent/mock"

	"github.com/golang/mock/gomock"
)

func call(ctx context.Context, m Math) (int, error) {
	result := make(chan int)
	go func() {
		result <- m.Sum(1, 2)
		close(result)
	}()
	select {
	case r := <-result:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// TestConcurrentFails is expected to fail (and is disabled). It
// demonstrates how to use gomock.WithContext to interrupt the test
// from a different goroutine.
func TestConcurrentFails(t *testing.T) {
	t.Skip("Test is expected to fail, remove skip to trying running yourself.")
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	m := mock.NewMockMath(ctrl)
	if _, err := call(ctx, m); err != nil {
		t.Error("call failed:", err)
	}
}

func TestConcurrentWorks(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	m := mock.NewMockMath(ctrl)
	m.EXPECT().Sum(1, 2).Return(3)
	if _, err := call(ctx, m); err != nil {
		t.Error("call failed:", err)
	}
}
