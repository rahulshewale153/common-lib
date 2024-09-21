package log

import (
	"bytes"
	"testing"
)

func TestLogger(t *testing.T) {
	// Create a buffer to capture the log output
	var buf bytes.Buffer
	logger := NewLogger()
	logger.SetOutput(&buf)
	logger.SetLogLevel(TRACE)

	// Test various log levels
	logger.Trace("This is a trace message.")
	logger.Debug("This is a debug message.")
	logger.Info("This is an info message.")
	logger.Warn("This is a warning message.")
	logger.Error("This is an error message.")
	logger.Critical("This is a critical message.")

	// Check if the output contains the log messages
	expected := []string{
		"[TRACE] This is a trace message.",
		"[DEBUG] This is a debug message.",
		"[INFO] This is an info message.",
		"[WARN] This is a warning message.",
		"[ERROR] This is an error message.",
		"[CRITICAL] This is a critical message.",
	}

	output := buf.String()
	for _, exp := range expected {
		if !contains(output, exp) {
			t.Errorf("Expected log message not found: %s", exp)
		}
	}
}

// Helper function to check if a substring exists in a string
func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
