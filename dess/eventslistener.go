package dess

import (
	"encoding/json"
	"log"

	"github.com/fsouza/go-dockerclient"
	"github.com/urfave/cli"
)

// Start is the alternative 'main' function of dess.
func Start(c *cli.Context) error {
	// First we set the global configuration
	setConfig(c)
	// This is a channel on which we can receive events
	events, err := events()
	if err != nil {
		return err
	}
	for {
		select {
		case msg := <-events:
			msgJSON, err := json.Marshal(msg)
			if err != nil {
				return err
			}
			msgJSONString := string(msgJSON)
			log.Printf("Received event: %v", msgJSONString)
			logSyslog(msgJSONString)
		}
	}
}

// events sregisters and starts an Event listener, and returns a channel
// that can be used to read the events from.
func events() (chan *docker.APIEvents, error) {
	var listener chan *docker.APIEvents
	var err error

	client, err := DockerClient()
	if err != nil {
		return listener, err
	}

	listener = make(chan *docker.APIEvents, 50)
	err = client.AddEventListener(listener)
	return listener, err
}
