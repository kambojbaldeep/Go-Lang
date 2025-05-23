/*Closures
A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

In this example, the concatter() function returns a function that has reference to an enclosed doc value. Each successive call to harryPotterAggregator mutates that same doc variable.

func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive
}

Assignment
Keeping track of how many texts we send is mission-critical at Textio. Complete the adder() enclosing function.

It should return a function that adds its input (an int) to an enclosed sum value, then return the new sum. In other words, it keeps a running total of the sum variable within a closure.*/

// The Solution is:

package main

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}