package tcp

import (
	"net"
	"fmt"
)

//a simple tcp server
func open() {

	addr, err := net.ResolveTCPAddr("tcp", ":10086")
	if err != nil {
		panic(err)
	}
	tcpListener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		fmt.Println(conn.RemoteAddr().String(), "---连接---")
		socket := Socket{
			Conn: conn,
		}
		go socket.Read()
	}

}

type SocketInterface interface {
	Write()
	Read()
}

type Socket struct {
	Conn *net.TCPConn
}

func (s *Socket) Write() {

}

func (s *Socket) Read() {
	for {
		b := make([]byte, 1024)
		n, err := s.Conn.Read(b)
		if err != nil {
			break
		}
		content := string(b[:n])
		fmt.Println(content)
		switch content {
		case "ping":
			s.Conn.Write([]byte("pong"))
			break
		default:
			s.Conn.Write([]byte("hello client"))
		}

	}
}
