[![Build Status](https://travis-ci.org/bjwschaap/docker-events-syslog.svg?branch=master)](https://travis-ci.org/bjwschaap/docker-events-syslog)
[![Go Report Card](https://goreportcard.com/badge/github.com/bjwschaap/docker-events-syslog)](https://goreportcard.com/report/github.com/bjwschaap/docker-events-syslog)

# Docker Events Syslog Streamer (dess)
Experimental daemon that forwards Docker events to the Syslog.

By default Docker events are not logged. In order to capture these events and
log them we use this `dess` tool. It logs each Docker event to the syslog in
JSON format.
