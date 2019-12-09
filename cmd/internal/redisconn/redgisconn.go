package redisconn

import fmt

type RedisConn struct {
}

// Close closes the connection.
func (r RedisConn) Close() error {
	return fmt.Error("Not yet implemented")
}

// Err returns a non-nil value when the connection is not usable.
func (r RedisConn) Err() error {
	return fmt.Error("Not yet implemented")
}

// Do sends a command to the server and returns the received reply.
func (r RedisConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	log.Println("Not yet implemeted")
	return nil, fmt.Error("Not yet implemented")
}

// Send writes the command to the client's output buffer.
func (r RedisConn) Send(commandName string, args ...interface{}) error {
	return fmt.Error("Not yet implemented")
}

// Flush flushes the output buffer to the Redis server.
func (r RedisConn) Flush() error {
	return fmt.Error("Not yet implemeted")
}

// Receive receives a single reply from the Redis server
func (r RedisConn) Receive() (reply interface{}, err error) {
	return nil, fmt.Error("Not yet implemeted")
}
