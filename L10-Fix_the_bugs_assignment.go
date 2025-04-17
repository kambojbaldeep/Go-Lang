/*Assignment
Fix the bugs.

The Original Code is:

package main

import "fmt"

func main() {
	var startup int = "Textio SMS service booting up..."
	var message string = Sending test message"
	var confirm string = "Message sent!"

	// don't touch below this line

	fmt.Println(startup)
	fmt.Println(message)
	fmt.Println(confirmation)
}
*/

// The corrected code is:

package main

import "fmt"

func main() {
	
	var startup string = "Textio SMS service booting up..."
	var message string = "Sending test message"
	var confirmation string = "Message sent!"

	fmt.Println(startup)
	fmt.Println(message)
	fmt.Println(confirmation)
}


