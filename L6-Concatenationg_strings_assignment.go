/* Assignment
Textio uses basic authentication to log users in.

The code on the right has a type error. Change the type of password to a string (but use the same numeric value) so that it can be concatenated with the username variable.

The code is: 
package main

import "fmt"

func main() {
	var username string = "presidentSkroob"
	var password int = 12345

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+password)
}

*/

// The Solution is:

package main

import "fmt"

func main() {
	var username string = "presidentSkroob"
	var password string = "12345"

	fmt.Println("Authorization: Basic", username+":"+password)
}