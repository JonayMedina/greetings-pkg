package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is required")
	}
	message := fmt.Sprintf("Hello, %v! desde greetings", name)
	return message, nil
}

func Hello2(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is required")
	}
	message := fmt.Sprintf("Hello, %v! desde greetings2", name)
	return message, nil
}
