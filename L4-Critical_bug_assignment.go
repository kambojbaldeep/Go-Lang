/* Assignment
Critical bug!

Textio users reported that we're billing them for wildly inaccurate amounts. They're supposed to be billed .02 dollars (2 cents) for each text message sent.

Without changing any of the other lines, fix the math bug on line 8.


The Oroginal code is:
package main

import "fmt"

func main() {
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage + float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}

*/


// The corrected Code is:

package main

import "fmt"

func main() {
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}