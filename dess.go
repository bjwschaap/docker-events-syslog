package main

import (
	"log"
	"os"

	"github.com/bjwschaap/docker-events-syslog/dess"
	"github.com/urfave/cli"
)

func main() {
	log.SetOutput(os.Stdout)

	app := cli.NewApp()
	app.Name = "dess"
	app.Usage = "Daemon that forwards Docker events to Splunk"
	app.Version = "0.1.0"
	app.Copyright = "(C)2016 Bastiaan Schaap"
	app.Author = "Bastiaan Schaap"
	app.Email = "bastiaan.schaap@gmail.com"
	app.UsageText = `./dess -s 123.123.1.2:4873 -u user -p pass
	 This daemon uses DOCKER_HOST, DOCKER_CERT_PATH and DOCKER_TLS settings to connect to the Docker daemon.
   Optionally you can override Docker settings with --host, --certs and --tls.`

	// Define the configuration flags the program can/should use
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host,H",
			Value:  "unix:///var/run/docker.sock",
			Usage:  "Address or location of the Docker API endpoint",
			EnvVar: "DOCKER_HOST",
		},
		cli.StringFlag{
			Name:   "certs,c",
			Value:  "",
			Usage:  "Location of TLS certificates for connecting with Docker",
			EnvVar: "DOCKER_CERT_PATH",
		},
		cli.BoolFlag{
			Name:   "tls",
			Usage:  "Should TLS be used for connecting to Docker",
			EnvVar: "DOCKER_TLS_VERIFY",
		},
		cli.StringFlag{
			Name:   "splunk,s",
			Value:  "",
			Usage:  "Address of Splunk",
			EnvVar: "SPLUNK_ADDRESS",
		},
		cli.StringFlag{
			Name:   "user,u",
			Value:  "",
			Usage:  "Username of the Splunk credentials",
			EnvVar: "SPLUNK_USER",
		},
		cli.StringFlag{
			Name:   "password,p",
			Value:  "",
			Usage:  "Password of the Splunk credentials",
			EnvVar: "SPLUNK_PASSWORD",
		},
	}

	// Set the main program logic
	app.Action = func(c *cli.Context) error {
		return dess.Start(c)
	}

	// Now start doing stuff
	app.Run(os.Args)
}
