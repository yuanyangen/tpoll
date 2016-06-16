package tpool

import (
	"net"
	"sync"
	"sync/atomic"
	"fmt"
)

// use another go routine to check the alive of tcp connnections
type root struct {
	pools     map[string]*pool
	searchMap map[*net.TCPConn]*pool
	mu        sync.Mutex
}

//a pool aims to a target
type pool struct {
	mu          sync.Mutex
	maxSize     int64
	currentSize int64
	identify    string
	conns       chan *net.TCPConn
}

var Root = &root{}

func init() {
	Root.pools = make(map[string]*pool)
	Root.searchMap = make(map[*net.TCPConn]*pool)
}

func GetTCPConn(address string) (*net.TCPConn, error) {
	return Root.getTCPConn(address)
}

func (p *pool) getTCPConn(addr string) (*net.TCPConn, error) {
	//if conn is not exist or  len < max or conn failed, try to dial a new conn
	if len(p.conns) <= 0 && p.currentSize < p.maxSize {
		tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
		if err != nil {
			return nil, err
		}

		// todo set timeout
		tcpConn, err := net.DialTCP("tcp4", nil, tcpAddr)
		if err != nil {
			return nil, err
		}
		atomic.AddInt64(&p.currentSize, 1)
		return tcpConn, nil
	}
	tcpConn := <-p.conns
	//todo check if the conn is alive
	return tcpConn, nil
}

func (r *root) getTCPConn(addr string) (*net.TCPConn, error) {
	r.mu.Lock()
	//the first conn
	if _, ok := r.pools[addr]; !ok {
		p := &pool{maxSize:10, identify:addr}
		p.conns = make(chan *net.TCPConn, 10)
		r.pools[addr] = p
	}
	r.mu.Unlock()
	pool := r.pools[addr]
	tcpConn, err := pool.getTCPConn(addr)
	if err == nil {
		r.mu.Lock()
		r.searchMap[tcpConn] = pool
		r.mu.Unlock()
	}
	return tcpConn, err
}

//if the conn is close
func PutTCPConn(tcpConn *net.TCPConn) {
	Root.putTCPConn(tcpConn)
}

func (r *root) putTCPConn(tcpConn *net.TCPConn) {
	pool := r.searchMap[tcpConn]
	pool.putTcpConn(tcpConn)
}

func (p *pool) putTcpConn(tcpConn *net.TCPConn) {
	p.conns <- tcpConn
}

