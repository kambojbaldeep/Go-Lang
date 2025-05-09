/* Assignment
Complete the reformat function. It takes a message string and a formatter function as input:

1. Apply the given formatter three times to the message
2. Add a prefix of TEXTIO: to the result
3. Return the final string
For example, if the message is "General Kenobi" and the formatter adds a period to the end of the string, the final result should be

TEXTIO: General Kenobi...  */

// The code is:

package main

func reformat(message string, formatter func(string) string) string {
	once := formatter(message)
	twice := formatter(once)
	thrice := formatter(twice)
	prefix := "TEXTIO: "
	return prefix + thrice
}