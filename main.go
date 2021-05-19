package main

import (
	"fmt"
	"math"
)

func main() {
	var bodyweight float64 = 192
	var heightInInches float64 = 70
	var weightInKg = bodyweight / 2.2
	var heightInCm = heightInInches * 2.54
	var age float64 = 30
	// Sedentary (little to no exercise + work a desk job) = 1.ve2
	// Lightly Active (light exercise 1-3 days / week) = 1.375
	// Moderately Active (moderate exercise 3-5 days / week) = 1.55
	// Very Active (heavy exercise 6-7 days / week) = 1.725
	// Extremely Active (very heavy exercise, hard labor job, training 2x / day) = 1.9
	var activityModifier = 1.55
	// Protein modifiers	:
	// Cut: 1.4 - 1.1 g/lb
	// Bulk: 0.8 - 1.0 g/lb
	var proteinModifier = 1.2
	// Fat modifiers:
	// Cut: 0.4 - 0.6 g/lb
	// Bulk: 20-30% calories
	var fatModifier = 0.5

	var BMR = 66 + (13.7 * weightInKg) + (1.8 * heightInCm) - (4.7 * age)
	var TDEE = math.Round(activityModifier * BMR)

	var proteinIntake = math.Round(bodyweight*proteinModifier*100) / 100
	var fatIntake = math.Round(bodyweight*fatModifier*100) / 100
	var calFromProtein = proteinIntake * 4
	var calFromFat = fatIntake * 9
	var calFromCarb = math.Round((TDEE-(calFromProtein+calFromFat))*100) / 100
	var carbIntake = calFromCarb / 4

	fmt.Println("Daily intake")
	fmt.Println("---")
	fmt.Println("TDEE:", TDEE, "cal/day")
	fmt.Println("Protein:", proteinIntake, "g,", calFromProtein, "cal")
	fmt.Println("Carbs:", carbIntake, "g,", calFromCarb, "cal")
	fmt.Println("Fat:", fatIntake, "g,", calFromFat, "cal")
}
