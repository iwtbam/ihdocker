package settings

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type IhSimpleLogFormatter struct {
}

func (s *IhSimpleLogFormatter) Format(entry *log.Entry) ([]byte, error) {

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var filename string
	var line int

	if entry.Caller != nil {
		filename = filepath.Base(entry.Caller.File)
		line = entry.Caller.Line
	}

	msg := fmt.Sprintf("[%v - %v - %v - %v] : %v\n", timestamp, filename, line, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}
