package dess

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

// DockerClient provides a connection object to use for interacting with
// the Docker API.
func DockerClient() (*docker.Client, error) {
	var client *docker.Client
	var err error

	if !config.dockerTLS {
		client, err = docker.NewClient(config.dockerHost)
	} else {
		ca := fmt.Sprintf("%s/ca.pem", config.dockerCertPath)
		cert := fmt.Sprintf("%s/cert.pem", config.dockerCertPath)
		key := fmt.Sprintf("%s/key.pem", config.dockerCertPath)
		client, err = docker.NewTLSClient(config.dockerHost, cert, key, ca)
	}
	return client, err
}
