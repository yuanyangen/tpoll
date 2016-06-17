package tpool

import (
	"net"
	"fmt"
)

func handel(){
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8090")
	tcpListner, _ :=  net.ListenTCP("tcp4", tcpAddr)

	for {
		tcpConn,_ := tcpListner.AcceptTCP()
		go do(tcpConn)
		//tcpConn.Close()
	}
}

func do(tcpConn *net.TCPConn) {
	fmt.Println("hahah connection ok")
	for {
		buf := make([]byte, 1024)
		tcpConn.Read(buf)
		tcpConn.Write(buf)
	}
}
