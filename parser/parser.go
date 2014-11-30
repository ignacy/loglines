package parser

import (
  "log"
  "regexp"

  "github.com/ignacy/loglines/types"
)

// Logline exaple:
// regexper b52bc8a Sun Oct 19 19:25:29 2014 +0200 having fun
const LOG_LINE_REGEXP = `([^\s]+)\s([^\s]+)\s(.*0100)\s(.*)`

func ParseLine(line string) *types.LogLine {
  re1, err := regexp.Compile(LOG_LINE_REGEXP)
  if err != nil {
    log.Fatal(err)
  }

  if result_slice := re1.FindAllStringSubmatch(line, -1); result_slice != nil {
    return &types.LogLine{
      Project: result_slice[0][1],
      Hash:    result_slice[0][2],
      Date:    result_slice[0][3],
      Message: result_slice[0][4],
    }
  } else {
    return &types.LogLine{}
  }
}
