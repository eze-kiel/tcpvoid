# TCPVoid

A simple TCP server for testing purposes. Supports HTTP, TCP Echo and TCP Void.

## Getting started

```
$ git clone github.com/eze-kiel/tcpvoid.git
$ cd tcpvoid/
$ make build
```

The binary can be found at `./out/bin/tcpvoid`.
Go is required.

## Usage

```
$ tcpvoid [KIND] [OPTIONS]

KINDS:
    http
        Start a HTTP server that return a given status code when joined
    echo
        Start a TCP server and echo anything sent on the socket
    void
        Start a TCP server that will absord anything and don't return anything

OPTIONS:
    -ip string
        Address to listen on (default "127.0.0.1")
    -p string
        Port to listen on (default "8080")
    -s int
        Status code to return with HTTP server (default 200)
```

## Examples

### HTTP

* Simplest example

```
$ tcpvoid http
INFO[0000] [HTTP] listening on 127.0.0.1:8080, returning 200
```

In another terminal:

```
$ curl -I 127.0.0.1:8080
HTTP/1.1 200 OK
Date: Fri, 23 Jul 2021 07:56:21 GMT
```

* Custom status code

```
$ tcpvoid -s 666 http
INFO[0000] [HTTP] listening on 127.0.0.1:8080, returning 666
```

```
$ curl -I 127.0.0.1:8080
HTTP/1.1 666 status code 666
Date: Fri, 23 Jul 2021 07:57:34 GMT
```

### Echo

```
$ tcpvoid echo
INFO[0000] [ECHO] listening on 127.0.0.1:8080
```

```
telnet 127.0.0.1 8080
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
hello there
hello there
Connection closed by foreign host.
```

### Void

```
$ tcpvoid void
INFO[0000] [VOID] listening on 127.0.0.1:8080 
```

```
$ telnet 127.0.0.1 8080
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
hello
there
...
hello ?
```

## License

MIT
