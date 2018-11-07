package log

import (
	"io"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

var receiver [512]byte

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		n, err := conn.Read(receiver[:])
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

	    log.Println(string(receiver[0:n]))
	}
}

func TestNewConnObject(t *testing.T) {
	// server
	server1 := func() {
		l, err := net.Listen(DefaultNetworkType, "127.0.0.1:2121")
		if err != nil {
			t.Errorf("Error listening: %v", err.Error())
			os.Exit(1)
		}
		defer l.Close()

		for {
			conn, _ := l.Accept()
			go handle(conn)
		}
	}

	server2 := func() {
		l, err := net.Listen(DefaultNetworkType, "127.0.0.1:2122")
		if err != nil {
			t.Errorf("Error listening: %v", err.Error())
			os.Exit(1)
		}
		defer l.Close()

		for {
			conn, _ := l.Accept()
			go handle(conn)
		}
	}

	go server1()
	go server2()

	// client
	co := NewConnObject(DefaultNetworkType, []string{"127.0.0.1:2121", "127.0.0.1:2122"})
	ticker := time.NewTicker(time.Duration(3) * time.Second)
	go func() {
		for {
			select {
			case <- ticker.C:
				for _, conn := range co.conns {
					co.Writing([]byte(conn.RemoteAddr().String()))
				}
			}
		}
	}()

	time.Sleep(time.Duration(30) * time.Second)
	ticker.Stop()
	co.Close()
}
