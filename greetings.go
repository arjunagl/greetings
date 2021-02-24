package greetings

import (
	"fmt"

	"github.com/arjunagl/greetings/v2/utility"

	log "github.com/sirupsen/logrus"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	log.WithFields(log.Fields{
		"animal": "simple logger 2.0.2",
	}).Info("A major logger version")

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	fmt.Println(utility.GiveNumber())
	return message
}
