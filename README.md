# tpoll

tpoll aims to provide tcp connection management , the pool is only necessary in client side


# install
```
go get github.com/yuanyangen/tpool
```

# usage

get a connection 
```
	tcpConn, _ := GetTCPConn("www.so.com:80")
	
```


put back this connection
```
	tcpConn, _ = GetTCPConn("www.so.com:80")

```
