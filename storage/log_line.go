package storage

type LogLine struct {
  Project string
  Hash    string
  Date    string
  Message string
}

func (line LogLine) String() string {
  return "Project: " + line.Project + " Message: " + line.Message + "\n"
}

func (line *LogLine) hasData() bool {
  return line.Project != "" && line.Message != ""
}
