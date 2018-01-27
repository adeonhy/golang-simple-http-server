package request

import (
  // "fmt"
  "strings"
)

type httpRequest struct {
  Method string
  Target string
  Version string
  Headers []string
  Body []byte
}

func Parse(buf []byte, n int) httpRequest {
  msg := strings.Split(string(buf[:n]), "\r\n")
  reqs := strings.Split(msg[0], " ") // "GET / HTTP/1.1"

  req := httpRequest{}
  req.Method = reqs[0]
  req.Target = reqs[1]
  req.Version = reqs[2]

  // fmt.Println(req)
  return req
}
