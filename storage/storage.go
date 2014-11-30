package storage

import (
  "bufio"
  "log"
  "os"
  "sync"

  "github.com/ignacy/loglines/parser"
  "github.com/ignacy/loglines/types"
)

func GetLogLines(path string) {
  inFile, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer inFile.Close()

  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines)

  results := make(chan *types.LogLine)

  var waitGroup sync.WaitGroup

  for scanner.Scan() {
    waitGroup.Add(1)
    go func(text string, results chan<- *types.LogLine) {
      if len(text) >= 0 {
        results <- parser.ParseLine(text)
      }
      waitGroup.Done()
    }(scanner.Text(), results)

  }

  go func() {
    waitGroup.Wait()
    close(results)
  }()

  display(results)
}

func display(results chan *types.LogLine) {
  // The channel blocks until a result is written to the channel.
  // Once the channel is closed the for loop terminates.
  for line := range results {
    if line.HasData() {
      log.Printf("Line: %s", line)
    }
  }
}
