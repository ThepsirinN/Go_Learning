package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// message := fmt.Sprintf("Hello %s, Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(name []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, val := range name {
		message, err := Hello(val)
		if err != nil {
			return nil, err
		}

		messages[val] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// slices of string (array must provide length or ...)
	formats := []string{
		"Hi, %s. Welcome!",
		"Great to see you, %s!",
		"Hail, %s! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
