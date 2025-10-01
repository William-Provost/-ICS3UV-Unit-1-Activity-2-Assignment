/*
	* Author: William Provost
	* Version 1.0.0
	* Date: 2025-09-30
	* This program solves question 1 of the math problems:
	* 1. Converting 34 Celsius to Fahrenheit
	*/

package main

import "fmt"

/*
This block converts 34 Celsius to Fahrenheit
Formula: F = 9 / 5 * C + 32
*/

func main() {
	fmt.Println("°C = 34        °F = ?") //prints out givens
	fmt.Println("Formula: F = 9 / 5 * C + 32") //prints out formula
	fmt.Println("34°C in Farenheit is: " + fmt.Sprint(9 / 5 * 34 + 32) + "°F") //does math and prints out sentence

	fmt.Println("\nDone.") // The program is done
}