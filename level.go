package log

import (
	"fmt"
	"strings"
)

// Log levels.
const (
	DebugLevel Level = iota
	InfoLevel
	NoticeLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	AlertLevel
	FatalLevel
)

// Level of the log
type Level uint8

// String returns the level as string
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case NoticeLevel:
		return "NOTICE"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	case AlertLevel:
		return "ALERT"
	case FatalLevel:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// ParseLevel parses a level
func ParseLevel(level string) (Level, error) {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return DebugLevel, nil
	case "INFO":
		return InfoLevel, nil
	case "NOTICE":
		return NoticeLevel, nil
	case "WARN":
		return WarnLevel, nil
	case "ERROR":
		return ErrorLevel, nil
	case "PANIC":
		return PanicLevel, nil
	case "ALERT":
		return AlertLevel, nil
	case "FATAL":
		return FatalLevel, nil
	default:
		return 255, fmt.Errorf("log: unknown level %q", level)
	}
}

// MarshalJSON implementation.
func (l Level) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", l)), nil
}

// UnmarshalJSON implementation.
func (l *Level) UnmarshalJSON(data []byte) error {
	level, err := ParseLevel(strings.Trim(string(data), "\""))
	if err != nil {
		return err
	}

	*l = level
	return nil
}
