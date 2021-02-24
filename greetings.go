package greetings

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	log.WithFields(log.Fields{
		"animal": "simple logger 1.0.1",
	}).Info("A major logger version")

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
