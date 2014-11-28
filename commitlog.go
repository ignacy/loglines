package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "regexp"
)

type LogLine struct {
  Project string
  Hash    string
  Date    string
  Message string
}

func main() {
  fmt.Println("HOEM")
  readLine("/Users/ignacymoryc/Dropbox/example-log")
}

func readLine(path string) {
  inFile, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer inFile.Close()

  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines)

  for scanner.Scan() {
    fmt.Printf("\nLog line: %+v", parseLine(scanner.Text()))
  }
}

func parseLine(line string) *LogLine {
  re1, err := regexp.Compile(`([^\s]+)\s([^\s]+)\s(.*0100)\s(.*)`)
  if err != nil {
    log.Fatal(err)
  }
  result_slice := re1.FindAllStringSubmatch(line, -1)
  return &LogLine{result_slice[0][1], result_slice[0][2],
    result_slice[0][3], result_slice[0][4]}
}
