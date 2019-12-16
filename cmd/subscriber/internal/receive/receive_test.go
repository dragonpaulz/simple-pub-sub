package receive_test

import (
	"fmt"
	"simple-pub-sub/cmd/subscriber/internal/receive"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testReceiver struct {
	queuedError error
	queuedNum   []int64
	NumChan     chan int
	DoneChan    chan error
}

func Init(numSent int) testReceiver {
	return testReceiver{
		queuedNum:   make([]int64, 0),
		NumChan:     make(chan int, numSent),
		DoneChan:    make(chan error),
		queuedError: nil,
	}
}

func (t testReceiver) Receive() interface{} {
	if len(t.queuedNum) > 0 {
		nextNum := t.queuedNum[0]
		t.queuedNum = t.queuedNum[1:]
		return nextNum
	} else if len(t.DoneChan) > 0 {
		err := t.queuedError
		t.queuedError = nil
		return err
	} else {
		time.Sleep(time.Millisecond)
		return nil
	}
}

func (t testReceiver) AddNumChan(nums ...int64) {
	t.queuedNum = append(t.queuedNum, nums...)
}

func TestReceiver_ReturnsNumbers_NumbersOnChannel(t *testing.T) {
	timeout := time.After(time.Second)

	fmt.Println("Starting test")
	tr := Init(3)
	fmt.Println("Done init")
	received := make([]interface{}, 3)
	msg1 := int64(1)
	msg2 := int64(2)
	msg3 := int64(3)
	fmt.Println("About to call Receive")
	go receive.Receive(tr.DoneChan, tr.NumChan, tr)
	go tr.AddNumChan(msg1, msg2, msg3)

	fmt.Println("Done with go calls")

	for i := 0; i < 2; i++ {
		select {
		case n := <-tr.NumChan:
			received[i] = n
			fmt.Println("Received ", n)
		case <-timeout:
			t.Fatal("Test timed out")
		}

	}

	rec1, ok1 := received[0].(int64)
	rec2, ok2 := received[1].(int64)
	rec3, ok3 := received[2].(int64)

	require.True(t, ok1)
	require.True(t, ok2)
	require.True(t, ok3)

	assert.Equal(t, msg1, rec1)
	assert.Equal(t, msg2, rec2)
	assert.Equal(t, msg3, rec3)
}
