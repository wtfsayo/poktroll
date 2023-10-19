package channel

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestObserver_Unsubscribe(t *testing.T) {
	var (
		onUnsubscribeCalled = false
		inputCh             = make(chan int, 1)
	)
	obsvr := &channelObserver[int]{
		observerMu: &sync.RWMutex{},
		// using a buffered channel to keep the test synchronous
		observerCh: inputCh,
		onUnsubscribe: func(toRemove *channelObserver[int]) {
			onUnsubscribeCalled = true
		},
	}

	// should initially be open
	require.Equal(t, false, obsvr.isClosed)

	inputCh <- 1
	require.Equal(t, false, obsvr.isClosed)

	obsvr.Unsubscribe()
	// should be isClosed after `#Unsubscribe()`
	require.Equal(t, true, obsvr.isClosed)
	require.True(t, onUnsubscribeCalled)
}

func TestObserver_ConcurrentUnsubscribe(t *testing.T) {
	var (
		onUnsubscribeCalled = false
		publishCh           = make(chan int, 1)
	)
	obsvr := &channelObserver[int]{
		ctx:        context.Background(),
		observerMu: &sync.RWMutex{},
		// using a buffered channel to keep the test synchronous
		observerCh: publishCh,
		onUnsubscribe: func(toRemove *channelObserver[int]) {
			onUnsubscribeCalled = true
		},
	}

	require.Equal(t, false, obsvr.isClosed, "observer channel should initially be open")

	// publish until the test cleanup runs
	done := make(chan struct{}, 1)
	go func() {
		for idx := 0; ; idx++ {
			// return when done receives; otherwise,
			select {
			case <-done:
				return
			default:
			}

			// publish a value
			obsvr.notify(idx)
		}
	}()
	// send on done when the test cleans up
	t.Cleanup(func() { done <- struct{}{} })

	// wait a bit, then assert that the observer is still open
	time.Sleep(10 * time.Millisecond)

	require.Equal(t, false, obsvr.isClosed)

	obsvr.Unsubscribe()
	// should be isClosed after `#Unsubscribe()`
	require.Equal(t, true, obsvr.isClosed)
	require.True(t, onUnsubscribeCalled)
}
