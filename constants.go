package logger

import "time"

const (
	TCP   = "tcp"
	UDP   = "udp"
	LOCAL = "local"
	FILE  = "file"
)

const MaxConnectorQ = 4

type LogLevel int

const (
	EMERG LogLevel = iota
	ALERT
	CRIT
	ERR
	WARN
	NOTIFY
	INFO
	DEBUG
	FORCE // force one of levels above
)

var LogLevels = map[LogLevel]string{
	EMERG:  "EMERGENCY",
	ALERT:  "ALERT",
	CRIT:   "CRITICAL",
	ERR:    "ERROR",
	WARN:   "WARNING",
	NOTIFY: "NOTIFICATION",
	INFO:   "INFO",
	DEBUG:  "DEBUG",
}

var LogLevelsInv = map[string]LogLevel{
	"EMERGENCY":    EMERG,
	"ALERT":        ALERT,
	"CRITICAL":     CRIT,
	"ERROR":        ERR,
	"WARNING":      WARN,
	"NOTIFICATION": NOTIFY,
	"INFO":         INFO,
	"DEBUG":        DEBUG,
}

const (
	TagStartUp  = "startup"
	TagConnect  = "connect"
	TagRoutine  = "routine"
	TagShutDown = "shutdown"
)

const (
	DelimiterV       = " | "
	DelimiterH       = " - "
	DelimiterA       = " * "
	defaultDelimiter = DelimiterV
)

var defaultReconnectionTime = time.Minute
