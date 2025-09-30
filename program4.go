/*
	* Author: William Provost
	* Version 1.0.0
	* Date: 2025-09-30
	* This program solves question 4 of the math problems:
 	* 4. Finding the gross pay of the employee Fred
	*/

// 4. This program calculates the gross pay for an employee named Fred, who worked 40 hours at an hourly wage of $8.20
// Formula: gross pay = hours worked * hourly wage

package main

import "fmt"

func main(){
	fmt.Println("hours worked = 40 hours        hourly wage = $8.20        gross pay = ?") //prints out givens
	fmt.Println("Formula: gross pay = hours worked * hourly wage") //prints out formula
	fmt.Println("The gross pay for Fred is: " + "$" + fmt.Sprint(40 * 8.20)) //does math and prints out sentence

	fmt.Println("\nDone.")
}