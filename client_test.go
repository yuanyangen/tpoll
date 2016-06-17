package tpool

import (
	"testing"
	"time"
)

func Test_GetTCPConn(t *testing.T) {
	tcpConn, _ := GetTCPConn("www.so.com:80")

	tcpConn.Write([]byte("dsfads"))
	buf := make([]byte, 1024)
	tcpConn.Read(buf)
	PutTCPConn(tcpConn)

	time.Sleep(1*time.Second)

	tcpConn, _ = GetTCPConn("www.so.com:80")
	tcpConn.Write([]byte("dsfads"))
	tcpConn.Read(buf)
	PutTCPConn(tcpConn)

}




