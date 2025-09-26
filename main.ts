/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-09-25
 * @fileoverview This program solves four math problems:
 * 1. Convert Celsius to Fahrenheit
 * 2. Area of a circle
 * 3. Volume of a sphere
 * 4. Gross pay calculation
 */

// 1. Convert Celsius to Fahrenheit
let celsius: number = 34;
let fahrenheit: number = (9 / 5) * celsius + 32;
console.log("1. Temperature Conversion:");
console.log("   34°C = " + fahrenheit + "°F");

// 2. Area of a circle with radius of 14
let radiusCircle: number = 14;
let areaCircle: number = Math.PI * radiusCircle * radiusCircle;
console.log("2. Area of Circle:");
console.log("   Radius = 14 cm, Area = " + areaCircle + " cm²");

// 3. Volume of a sphere with radius of 4
let radiusSphere: number = 4;
let volumeSphere: number = (4 / 3) * Math.PI * Math.pow(radiusSphere, 3);
console.log("3. Volume of Sphere:");
console.log("   Radius = 4 cm, Volume = " + volumeSphere + " cm³");

// 4. Freds pay
let nameEmployee: string = "Fred";
let hoursWorked: number = 40;
let hourlyWage: number = 8.20;
let grossPay: number = hoursWorked * hourlyWage;
console.log("4. Gross Pay:");
console.log(
  "   Employee: " +
  nameEmployee +
  ", Hours: " +
  hoursWorked +
  ", Wage: $" +
  hourlyWage +
  ", Gross Pay = $" +
  grossPay.toFixed(2)
);
