// The godemo package exports a toy "Greet" function that we use to demonstrate Unit Testing.
package hello

import "fmt"

// Greet function returns a friendly greeting.
func Greet(name string) string {
	if name == "" {
		name = "Diego"
	}
	return fmt.Sprintf("Howdy %s!", name)
}
