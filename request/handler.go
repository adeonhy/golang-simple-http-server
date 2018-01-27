package request

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
  "os"
  "path/filepath"
)

type httpResponse struct {
	Code int
	Type string
	Body []byte
}

var BadRequest = httpResponse{Code: 400,
	Type: "text/html",
	Body: readAsset("./public/400.html")}

var Forbidden = httpResponse{Code: 403,
	Type: "text/html",
	Body: readAsset("./public/403.html")}

var NotFound = httpResponse{Code: 404,
	Type: "text/html",
	Body: readAsset("./public/404.html")}

var ServerError = httpResponse{Code: 500,
	Type: "text/html",
	Body: readAsset("./public/500.html")}

func responseMsg(res httpResponse) []byte {
	msgLines := []string{fmt.Sprintf("HTTP/1.1 %d", res.Code),
		"Date: " + time.Now().String(),
		"Server: \"Golang SimpleHTTPServer/0.1\"",
		"Content-Type: " + res.Type,
		"Content-Length: " + fmt.Sprint(len(res.Body)),
		"Connection: Close",
		"",
		""}

	resData := append([]byte(strings.Join(msgLines, "\r\n")), res.Body...)

	return resData
}

func readAsset(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("read asset", err)
		return nil
	}
	return data
}

func Handle(req httpRequest) []byte {
  switch req.Method {
  case "GET":
    return handlePath(req.Target)
  default:
    return responseMsg(BadRequest)
  }
}

func handlePath(path string) []byte {
  exe, err := os.Executable()
  if err != nil {
    log.Println("os.Executable", err)
    return responseMsg(ServerError)
  }

  exePath := filepath.Dir(exe)
  serverRoot := filepath.Join(exePath, "public")
  targetPath := filepath.Join(serverRoot, path)

  fInfo, err := os.Stat(targetPath)
  if err != nil {
    if os.IsNotExist(err) {
      return responseMsg(NotFound)
    } else if os.IsPermission(err) {
      return responseMsg(Forbidden)
    } else {
      return responseMsg(ServerError)
    }
  }

  if fInfo.IsDir() {
    return handlePath(filepath.Join(path, "index.html"))
  }

  res := httpResponse{Code: 200, 
    Type: Mime(targetPath),
    Body: readAsset(targetPath)}

  return responseMsg(res)
}
