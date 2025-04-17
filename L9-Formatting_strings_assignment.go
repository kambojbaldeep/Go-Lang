/* Assignment
Create a new variable called msg on line 9 and use the appropriate formatting function to return a string that contains following:

Hi NAME, your open rate is OPENRATE percentNEWLINE

Replace NAME with the variable name,
Replace OPENRATE with the variable openRate rounded to the nearest "tenths" place, e.g 10.523 should be rounded down to 10.5
The word percent should appear as part of the string following the open rate value
Replace NEWLINE with the newline \n escape sequence 

Original Code:

package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	// don't edit below this line

	fmt.Print(msg)
} 
*/

// The corrected code is:

package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)

	fmt.Print(msg)
}