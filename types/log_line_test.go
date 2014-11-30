package types

import (
  "testing"
)

func TestHasData(t *testing.T) {
  logLine := LogLine{
    Project: "hello",
    Message: "heja",
  }

  if logLine.HasData() != true {
    t.Error("Failed to recognize LogLine has data")
  }
}
