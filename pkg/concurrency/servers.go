package concurrency

import "net"

func simplehandler(c net.Conn) {
	c.Write([]byte("ok"))
	c.Close()
}

func handler(c net.Conn, ch chan string) {
	ch <- c.RemoteAddr().String()
	c.Write([]byte("ok"))
	c.Close()
}

func DoServers() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go simplehandler(c)
	}
}
