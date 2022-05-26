// Golang program to illustrate the
// use of strings.Join Function

package main

// importing fmt and strings
import (
	"fmt"
	"strings"
)

// calling main method
func main() {
	// array of strings.
	str := []string{"A", "Computer-science", "portal", "for", "Geeks"}
	// joining the string by separator in middle.
	fmt.Println(strings.Join(str, " "))
}
