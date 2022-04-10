package redisprotocol

import (
	"bufio"
	"fmt"
	"net"
)

// SimpleReidsPotocol redis protocol struct
type SimpleReidsPotocol struct {
	conn net.Conn
}

// NewSimpleReidsPotocol create redis protocol instance
func NewSimpleReidsPotocol(conn net.Conn) *SimpleReidsPotocol {
	return &SimpleReidsPotocol{conn: conn}
}

// Send send a raw redis command
func (srp *SimpleReidsPotocol) Send(cmd ICmd) {
	cmdEncoder := NewCmdEncoder()
	msg := cmdEncoder.Encode(cmd)
	fmt.Println("send raw command is", msg)
	fmt.Fprint(srp.conn, msg)
}

// Receive receive a raw redis  command
func (srp *SimpleReidsPotocol) Receive() {
	scanner := bufio.NewScanner(srp.conn)
	for {
		if ok := scanner.Scan(); !ok {
			break
		}
		fmt.Println("receive message is", scanner.Text())
	}
	fmt.Println("receive empty")
}
