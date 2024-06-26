package notify

import (
	"runtime"
	"strings"
)

const (
	SeverityLow = iota
	SeverityNormal
	SeverityUrgent
)

type Severity int

type Notify struct {
	title    string
	message  string
	severity Severity
}

func New(title, message string, severity Severity) *Notify {
	return &Notify{
		title:    title,
		message:  message,
		severity: severity,
	}
}

// String method returns the correct string for the severity for each OS
func (s Severity) String() string {
	sev := "low"

	switch s {
	case SeverityLow:
		sev = "low"
	case SeverityNormal:
		sev = "normal"
	case SeverityUrgent:
		sev = "critical"
	}

	if runtime.GOOS == "darwin" {
		sev = strings.Title(sev)
	}

	if runtime.GOOS == "windows" {
		switch s {
		case SeverityLow:
			sev = "Info"
		case SeverityNormal:
			sev = "Warning"
		case SeverityUrgent:
			sev = "Error"
		}
	}

	return sev
}
