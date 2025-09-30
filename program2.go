/*
	* Author: William Provost
	* Version 1.0.0
	* Date: 2025-09-30
	* This program solves question 2 of the math problems:
 	* 2. Finding the area of a circle
	*/

// 2. This program calculates the area of a circle with a radius of 14 cm
// Formula: A = Pi * r * r

package main

import "fmt"

func main(){
	fmt.Println("r = 14 cm        A = ?") //prints out givens
	fmt.Println("Formula: A = Pi * r * r") //prints out formula
	fmt.Println("The area of a circle with a radius of 14 cm is: " + fmt.Sprint(3.14 * 14 * 14) + "cm²") //does math and prints out sentence

	fmt.Println("\nDone.")
}