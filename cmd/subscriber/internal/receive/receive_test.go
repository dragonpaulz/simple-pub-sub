package receive_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"simple-pub-sub/cmd/subscriber/internal/receive"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockReceiver struct {
	mock.Mock
}

func (m *mockReceiver) Receive() interface{} {
	args := m.Called()
	return args.Get(0)
}

func TestReceive_mockReceiver(t *testing.T) {
	timeout := time.After(100 * time.Millisecond)
	testobj := new(mockReceiver)

	testobj.On("Receive").Return(redis.Message{Data: []byte(strconv.Itoa(1))})
	NumChan := make(chan int)
	DoneChan := make(chan error)

	go receive.Receive(DoneChan, NumChan, testobj)

	var received int

	select {
	case n := <-NumChan:
		received = n
		fmt.Println("Received ", n)
	case <-timeout:
		t.Fatal("Test timed out")
	}

	assert.Equal(t, 1, received)
}

func TestReceiveMultiple_mockReceiver(t *testing.T) {
	timeout := time.After(100 * time.Millisecond)
	testobj := new(mockReceiver)

	testobj.On("Receive").Return(redis.Message{Data: []byte(strconv.Itoa(1))}).Once()
	testobj.On("Receive").Return(redis.Message{Data: []byte(strconv.Itoa(2))})
	NumChan := make(chan int)
	DoneChan := make(chan error)

	go receive.Receive(DoneChan, NumChan, testobj)

	received := make([]int, 2)

	for i := 0; i < 2; i++ {
		select {
		case n := <-NumChan:
			received[i] = n
			fmt.Println("Received ", n)
		case <-timeout:
			t.Fatal("Test timed out")
		}
	}

	assert.Contains(t, received, 1)
	assert.Contains(t, received, 2)
}
