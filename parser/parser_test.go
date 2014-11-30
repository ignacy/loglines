package parser

import "testing"

func TestParsingLogLine(t *testing.T) {
  line := "some_project b52bc8a Sun Oct 19 19:25:29 2014 +0200 stuff"
  logLine := ParseLine(line)

  if logLine.Project != "some_project" {
    t.Error("Failed to parse project name correctly")
  }

  if logLine.Message != "stuff" {
    t.Error("Failed to parse commit message correctly")
  }
}
