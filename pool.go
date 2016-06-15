package tpoll

import (
	"net"
	"sync"
)

//
type root struct {
	pools     map[string]*pool
	searchMap map[*net.TCPConn]*pool
	mu        sync.Mutex
}


//a pool aims to a target
type pool struct {
	maxSize     int
	currentSize int
	identify    string
	conns       chan *net.TCPConn
}

var root = &root{}

func init() {

}

func GetTCPConn(ip string, port int) *net.TCPConn {

	return root.getTCPConn(ip, string(port))
}

//if the conn is close
func PutTCPConn(*net.TCPConn) {

}

func (p *pool) getTcpConn(ip string, port string) *net.TCPConn {
	//if conn is not exist or  len < max or conn failed, try to dial a new conn
	if len(p.conns) <= 0 && p.currentSize < p.maxSize {
		conn := net.DialTCP("tcp4")


	}

	//check if the conn is alive

	return
}

func (r *root) getTCPConn(ip string, port string) *net.TCPConn {
	key := ip + port
	r.mu.Lock()
	//the first conn
	if _, ok := r.pools[key]; !ok {
		r.pools[key] = &pool{maxSize:10, identify:key}
	}
	r.mu.Unlock()
	pool := r.pools[key]
	return pool.getTcpConn(ip, port)
}

func (r *root) getPools() {}


