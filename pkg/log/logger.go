package log

import (
	"fmt"
	"log"
	"time"
)

type LogWriter struct {
	level string
}

func (lw *LogWriter) Write(bytes []byte) (int, error) {
	message := fmt.Sprintf("[%s][%s] - %s", time.Now().Format(time.RFC3339Nano), lw.level, string(bytes))
	return fmt.Print(message)
}

func (lw *LogWriter) Init(level string) {
	lw.level = "DEBUG"
	log.SetFlags(0)
	log.SetOutput(lw)
}
