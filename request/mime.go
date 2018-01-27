package request

import (
  "path/filepath"
)

var mimeTypes = map[string]string{".html": "text/html",
  ".jpg": "image/jpeg",
  ".png": "image/png"}

func Mime(path string) string {
  ext := filepath.Ext(path)
  mime, ok := mimeTypes[ext]

  if ok {
    return mime
  } else {
    return "application/octet-stream"
  }
}
