package storage

import (
  "bufio"
  "log"
  "os"
  "regexp"
  "sync"
)

const LOG_LINE_REGEXP = `([^\s]+)\s([^\s]+)\s(.*0100)\s(.*)`

func GetLogLines(path string) {
  inFile, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer inFile.Close()

  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines)

  results := make(chan *LogLine)

  var waitGroup sync.WaitGroup

  for scanner.Scan() {
    waitGroup.Add(1)
    go func(text string, results chan<- *LogLine) {
      if len(text) >= 0 {
        results <- parseLine(text)
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

func parseLine(line string) *LogLine {
  re1, err := regexp.Compile(LOG_LINE_REGEXP)
  if err != nil {
    log.Fatal(err)
  }

  if result_slice := re1.FindAllStringSubmatch(line, -1); result_slice != nil {
    return &LogLine{
      Project: result_slice[0][1],
      Hash:    result_slice[0][2],
      Date:    result_slice[0][3],
      Message: result_slice[0][4],
    }
  } else {
    return &LogLine{}
  }
}

func display(results chan *LogLine) {
  // The channel blocks until a result is written to the channel.
  // Once the channel is closed the for loop terminates.
  for line := range results {
    if line.hasData() {
      log.Printf("Line: %s", line)
    }
  }
}
