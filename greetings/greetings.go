package greetings
import "fmt"

// Hello prints a greeting for the named person
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
