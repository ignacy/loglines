package main

import (
  "log"
  "os"

  "github.com/ignacy/loglines/storage"

  _ "github.com/ignacy/loglines/types"
)

func init() {
  log.SetFlags(0)
  log.SetOutput(os.Stdout)
}

func main() {
  if len(os.Args) < 2 {
    log.Fatal("Usege: loglines PATH_TO_LOGLINES_FILE")
  }
  storage.GetLogLines(os.Args[1])
}
