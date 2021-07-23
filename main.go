package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/sirupsen/logrus"
)

type app struct {
	address    string
	port       string
	statusCode int
}

func main() {
	var a app
	flag.StringVar(&a.address, "ip", "127.0.0.1", "Address to listen on")
	flag.StringVar(&a.port, "p", "8080", "Port to listen on")
	flag.IntVar(&a.statusCode, "s", 200, "Status code to return with HTTP server")
	flag.Parse()

	if len(flag.Args()) == 0 {
		logrus.Fatal("no kind specified")
	} else if len(flag.Args()) > 1 {
		logrus.Fatal("only one kind can be specified")
	}

	kind := flag.Arg(0)
	switch kind {
	case "http":
		logrus.Info(fmt.Sprintf("[HTTP] listening on %s:%s, returning %d", a.address, a.port, a.statusCode))
		if err := a.serveHttp(); err != nil {
			logrus.Fatal(err)
		}
	case "echo":
		logrus.Info(fmt.Sprintf("[ECHO] listening on %s:%s", a.address, a.port))
		if err := a.serveEcho(); err != nil {
			logrus.Fatal(err)
		}
	case "void":
		logrus.Info(fmt.Sprintf("[VOID] listening on %s:%s", a.address, a.port))
		if err := a.serveVoid(); err != nil {
			logrus.Fatal(err)
		}
	default:
		logrus.Fatal(fmt.Sprintf("unknown kind: %s", kind))
	}
}

func (a *app) serveHttp() error {
	srv := http.Server{
		Addr: a.address + ":" + a.port,
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(a.statusCode)
		}),
	}
	return srv.ListenAndServe()
}

func (a *app) serveEcho() error {
	l, err := net.Listen("tcp", a.address+":"+a.port)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go func(c net.Conn) {
			defer c.Close()

			data, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				logrus.Fatal(err)
			}

			_, err = c.Write([]byte(data))
			if err != nil {
				logrus.Fatal(err)
			}
		}(conn)
	}
}

func (a *app) serveVoid() error {
	l, err := net.Listen("tcp", a.address+":"+a.port)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go func(c net.Conn) {
			defer c.Close()
			select {}
		}(conn)
	}
}
