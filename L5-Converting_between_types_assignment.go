/* Assignment
Our Textio customers want to know how long they have had accounts with us.

On line 7, create a accountAgeInt variable and assign it the value of accountAgeFloat truncated to an integer.
*/

package main

import "fmt"

func main() {
	accountAgeFloat := 2.6

	accountAgeInt := int(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
