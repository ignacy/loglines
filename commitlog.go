package main

import (
  "bufio"
  "log"
  "os"
  "regexp"
  "sync"
)

type LogLine struct {
  Project string
  Hash    string
  Date    string
  Message string
}

func main() {
  //readLine("/Users/ignacymoryc/Dropbox/example-log")
  readLine("/Users/ignacymoryc/Dropbox/git-commit-logs")
}

func readLine(path string) {
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
      results <- parseLine(text)
      waitGroup.Done()
    }(scanner.Text(), results)

  }

  go func() {
    waitGroup.Wait()
    close(results)
  }()

  Display(results)
}

func parseLine(line string) *LogLine {
  re1, err := regexp.Compile(`([^\s]+)\s([^\s]+)\s(.*0100)\s(.*)`)
  if err != nil {
    log.Fatal(err)
  }
  result_slice := re1.FindAllStringSubmatch(line, -1)
  if result_slice == nil {
    return &LogLine{}
  } else {
    return &LogLine{result_slice[0][1], result_slice[0][2],
      result_slice[0][3], result_slice[0][4]}
  }
}

func Display(results chan *LogLine) {
  // The channel blocks until a result is written to the channel.
  // Once the channel is closed the for loop terminates.
  for line := range results {
    log.Printf("%s:\n%s\n", line.Project, line.Message)
  }
}
