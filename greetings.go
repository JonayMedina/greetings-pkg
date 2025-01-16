package greetings

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %v! desde greetings", name)
}

func Hello2(name string) string {
	return fmt.Sprintf("Hello, %v! desde greetings2", name)
}

