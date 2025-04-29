/* Assignment
Fix the bugs in the monthlyBillIncrease and getBillForMonth functions.

monthlyBillIncrease: Returns the increase in the bill from the previous to the current month. If the bill decreased, it should return a negative number.
getBillForMonth: Returns the bill for the given month.
It looks like whoever wrote the getBillForMonth function thought that they could pass in the bill parameter, update it inside the function, and that update would apply in the parent function (monthlyBillIncrease). They were wrong.

Change the getBillForMonth function to explicitly return the bill for the given month, and be sure to capture that return value properly in the monthlyBillIncrease function.

The function signature for getBillForMonth should only take 2 parameters once you're done. */

// The correct ans is:

package main

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill := getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}
