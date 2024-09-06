package logger

import (
	"bytes"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	l := NewLogger(&buf, INFO)
	l.Info("info message")
	l.Warn("warn message")
	l.Error("error message")
	l.Debug("debug message")
	got := buf.String()
	timeStamp:= time.Now().Format("2006-01-02 15:04:05")
	want:= timeStamp + " [INFO] info message\n" + timeStamp + " [WARN] warn message\n" + timeStamp + " [ERROR] error message\n" + timeStamp + " [DEBUG] debug message\n"
	if got != want {
		t.Errorf("\ngot- %q\n want- %q", got, want)
	}
}
