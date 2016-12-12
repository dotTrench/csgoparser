package receiver

import (
	"bufio"
	"net"
)

type UDPReceiever struct {
	conn   *net.UDPConn
	reader bufio.Reader
}

func (r *UDPReceiever) GetLine() {

}
