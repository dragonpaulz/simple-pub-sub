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
