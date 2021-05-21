// 192 70 30 1.55
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func roundToTwoSigFigs(i float64) float64 {
	return math.Round(i*100) / 100
}

func main() {
	/*
		Until flags are implemented - you need to run program with args in location:
		1: Bodyweight
		2: Height (inches)
		3. Age
		4. Activity level (1-5)
	*/
	var err error
	var bodyweight float64
	var heightInInches float64
	var age float64

	var activityModifier float64
	for idx, arg := range os.Args[1:] {
		var activityModifierInput float64
		switch idx {
		case 0:
			bodyweight, err = strconv.ParseFloat(arg, 64)
		case 1:
			heightInInches, err = strconv.ParseFloat(arg, 64)
		case 2:
			age, err = strconv.ParseFloat(arg, 64)
		case 3:
			activityModifierInput, err = strconv.ParseFloat(arg, 64)
			if activityModifierInput == 1 {
				// 1. Sedentary (little to no exercise + work a desk job) = 1.2
				activityModifier = 1.2
			} else if activityModifierInput == 2 {
				// 2. Lightly Active (light exercise 1-3 days / week) = 1.375
				activityModifier = 1.55
			} else if activityModifierInput == 3 {
				// 3. Moderately Active (moderate exercise 3-5 days / week) = 1.55
				activityModifier = 1.55
			} else if activityModifierInput == 4 {
				// 4. Very Active (heavy exercise 6-7 days / week) = 1.725
				activityModifier = 1.725
			} else if activityModifierInput == 5 {
				// 5. Extremely Active (very heavy exercise, hard labor job, training 2x / day) = 1.9
				activityModifier = 1.9
			}
			activityModifier, err = strconv.ParseFloat(arg, 64)
		}
		if err != nil {
			fmt.Println("Error", err)
		}
	}
	var weightInKg = bodyweight / 2.2
	var heightInCm = heightInInches * 2.54

	// Protein modifiers:
	// Cut: 1.4 - 1.1 g/lb
	// Bulk: 0.8 - 1.0 g/lb
	var proteinModifier = 1.2
	// Fat modifiers:
	// Cut: 0.4 - 0.6 g/lb
	// Bulk: 20-30% calories
	var fatModifier = 0.5

	var BMR = 66 + (13.7 * weightInKg) + (1.8 * heightInCm) - (4.7 * age)
	var TDEE = math.Round(activityModifier * BMR)

	var proteinIntake = roundToTwoSigFigs(bodyweight * proteinModifier)
	var fatIntake = roundToTwoSigFigs(bodyweight * fatModifier)
	var calFromProtein = proteinIntake * 4
	var calFromFat = fatIntake * 9
	var calFromCarb = roundToTwoSigFigs((TDEE - (calFromProtein + calFromFat)))
	var carbIntake = calFromCarb / 4

	fmt.Println("Daily intake")
	fmt.Println("---")
	fmt.Println("TDEE:", TDEE, "cal/day")
	fmt.Println("Protein:", proteinIntake, "g,", calFromProtein, "cal")
	fmt.Println("Carbs:", carbIntake, "g,", calFromCarb, "cal")
	fmt.Println("Fat:", fatIntake, "g,", calFromFat, "cal")
}
