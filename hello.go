// The godemo package exports interesting functions for having a great Tuesday
// afternoon.
package godemo

import "fmt"

// Greet function returns a friendly greeting.
func Greet(name string) string {
	if name == "" {
		name = "Diego"
	}
	return fmt.Sprintf("Howdy %s!", name)
}
