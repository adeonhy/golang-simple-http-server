package main

import (
  "log"
  "net"
  "strings"

  "./request"
)

func startServer(addrString string)(server *net.TCPListener, err error) {
  addr, err := net.ResolveTCPAddr("tcp", addrString)
  if err != nil {
    log.Println("error on ResolveTCPAddr:", err)
    return nil, err
  }

  server, err = net.ListenTCP("tcp", addr)
  if err != nil {
    log.Println("error on ListenTCP:", err)
    return nil, err
  }
  log.Println("server listening on ", server)

  return server, nil
}

func handleListener(server *net.TCPListener) error {
  defer server.Close()

  for {
    conn, err := server.AcceptTCP()
    if err != nil {
      log.Println("acceptTCP", err)
      return err
    }
    log.Println("accepted on ", conn)

    go handleConnection(conn)
  }
}

func handleConnection(conn *net.TCPConn) {
  defer conn.Close()

  buf := make([]byte, 1024*1024)

  for {
    n, err := conn.Read(buf)
    if err != nil {
      log.Println("error on Read:", err)
      return
    }

    log.Println(buf[:n])
    log.Println(strings.TrimRight(string(buf[:n]), "\r\n"))

    req := request.Parse(buf, n)
    log.Println("request: ", req)
    res := request.Handle(req)

    conn.Write([]byte(res))
  }
}

func main() {
  server, err := startServer("0.0.0.0:8080")
  if err != nil {
    log.Println("startServer", err)
    return
  }

  handleListener(server)
}
