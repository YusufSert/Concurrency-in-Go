package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)


func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{} // connection object
}

func warmServiceConnCache() *sync.Pool {
	p := sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}

	return &p
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()
	

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatal(err)
		}
		defer server.Close()
	
		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Printf("cannot accept connection %v", err)
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "200-ok")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}






func init() {
	demonStarted := startNetworkDaemon()
	demonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < 10; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		buf :=  make([]byte, 1024)
		if buf, err = io.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		fmt.Printf("%b\n", buf)
	
		conn.Close()
	}
}