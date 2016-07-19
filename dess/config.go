package dess

import (
	"log"

	"github.com/urfave/cli"
)

// Configuration is a structure to hold global configuration settings
type configuration struct {
	dockerHost     string
	dockerTLS      bool
	dockerCertPath string
	splunkAddress  string
	splunkUser     string
	splunkPassword string
}

// Global configuration for this service
var config configuration

// setConfig copies cli and environment parameters/settings/variables to
// the global configuration
func setConfig(c *cli.Context) {
	config.dockerHost = c.String("host")
	config.dockerTLS = c.Bool("tls")
	config.dockerCertPath = c.String("certs")
	config.splunkAddress = c.String("splunk")
	config.splunkUser = c.String("user")
	config.splunkPassword = c.String("password")
	passwordIsSet := config.splunkPassword != ""

	log.Printf("Docker host             : %v", config.dockerHost)
	log.Printf("Docker use TLS          : %v", config.dockerTLS)
	log.Printf("Docker TLS certificates : %v", config.dockerCertPath)
	log.Printf("Splunk address          : %v", config.splunkAddress)
	log.Printf("Splunk user             : %v", config.splunkUser)
	log.Printf("Splunk password is set  : %v", passwordIsSet)
}
