package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

func Hellos(names []string) (map[string]string, error) {
	// make(map[key-type]value-type)
	// key-value pairs (map) with key as string and value as string
	msgs := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name)

		if err != nil {
			return nil, err
		}

		msgs[name] = msg
	}
	return msgs, nil
}

// Go executes init functions automatically at program startup, after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
