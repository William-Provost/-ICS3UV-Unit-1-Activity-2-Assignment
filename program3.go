/*
	* Author: William Provost
	* Version 1.0.0
	* Date: 2025-09-30
	* This program solves question 3 of the math problems:
 	* 3. Finding the volume of a sphere
	*/

// 3. This program outputs the volume of a sphere with a radius of 4 cm
// Formula: V = 4 / 3 * Pi * r * r * r

package main

import "fmt"

func main(){
	fmt.Println("r = 4 cm        V = ?") //prints out givens
	fmt.Println("Formula: V = 4 / 3 * Pi * r * r * r") //prints out formula
	fmt.Println("The volume of a sphere with a radius of 4 cm is: " + fmt.Sprint(4 / 3 * 3.14 * 4 * 4 * 4) + "cm³") //does math and prints out sentence

	fmt.Println("\nDone.")
}