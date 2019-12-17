package receive_test

import (
	"container/list"
	"fmt"
	"simple-pub-sub/cmd/subscriber/internal/receive"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testReceiver struct {
	queuedError error
	mux         sync.Mutex
	queuedNum   *list.List
	NumChan     chan int
	DoneChan    chan error
}

func Init(numSent int) testReceiver {
	return testReceiver{
		queuedNum:   list.New(),
		NumChan:     make(chan int, numSent),
		DoneChan:    make(chan error),
		queuedError: nil,
	}
}

func (t testReceiver) Receive() interface{} {
	if t.queuedNum.Len() > 0 {
		nextNum := t.queuedNum.Front()
		var data []byte
		if i, ok := nextNum.Value.(int); ok {
			fmt.Println("i: ", i)
			s := strconv.Itoa(i)
			fmt.Println("s: ", s)
			data = []byte(s)
			fmt.Println("data: ", data)
		}

		t.queuedNum.Remove(nextNum)

		return redis.Message{Channel: "", Pattern: "", Data: data}
		// } else if t.queuedNum.Len() == 0 {
		// 	return fmt.Errorf("Done")
	} else if len(t.DoneChan) > 0 {
		fmt.Println("Done")
		err := t.queuedError
		t.queuedError = nil
		return err
	} else {
		time.Sleep(50 * time.Millisecond)
		return nil
	}
}

func (t testReceiver) AddNumChan(nums ...int) {
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, n := range nums {
		t.queuedNum.PushBack(n)
	}
}

func TestReceiver_ReturnsNumbers_NumbersOnChannel(t *testing.T) {
	timeout := time.After(time.Second)
	tr := Init(3)
	received := make([]interface{}, 3)
	msg1 := 1
	msg2 := 2
	msg3 := 3
	go receive.Receive(tr.DoneChan, tr.NumChan, tr)
	go tr.AddNumChan(msg1, msg2, msg3)

	for i := 0; i <= 2; i++ {
		fmt.Println("Run ", i)
		fmt.Printf("NumChan: %p\n", &tr.NumChan)
		select {
		case n := <-tr.NumChan:
			received[i] = n
			fmt.Println("Received ", n)
		case <-timeout:
			t.Fatal("Test timed out")
		}

	}

	fmt.Println("out")

	// This should probably be something else
	tr.DoneChan <- fmt.Errorf("Done test")

	rec1, ok1 := received[0].(int)
	rec2, ok2 := received[1].(int)
	rec3, ok3 := received[2].(int)

	require.True(t, ok1)
	require.True(t, ok2)
	require.True(t, ok3)

	assert.Equal(t, msg1, rec1)
	assert.Equal(t, msg2, rec2)
	assert.Equal(t, msg3, rec3)
}
