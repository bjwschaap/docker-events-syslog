package dess

import (
	"log"
	"log/syslog"
)

var slog *syslog.Writer

func init() {
	var err error
	slog, err = syslog.New(syslog.LOG_INFO, "DockerEvent")
	defer slog.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func logSyslog(msg string) {
	slog.Notice(msg)
}
