package main

import (
	"flag"
	"fmt"
	"math"
)

func roundToTwoSigFigs(i float64) float64 {
	return math.Round(i*100) / 100
}

func printOutput(TDEE float64, proteinIntake float64, calFromProtein float64, carbIntake float64, calFromCarb float64, fatIntake float64, calFromFat float64) {
	fmt.Println("Daily intake")
	fmt.Println("---")
	fmt.Println("TDEE:", TDEE, "cal/day")
	fmt.Println("Protein:", proteinIntake, "g,", calFromProtein, "cal")
	fmt.Println("Carbs:", carbIntake, "g,", calFromCarb, "cal")
	fmt.Println("Fat:", fatIntake, "g,", calFromFat, "cal")
}

func convertFlagsToVars() (float64, float64, float64, float64) {
	bodyweight := flag.Int("bw", 0, "bodyweight in lbs")
	height := flag.Int("h", 0, "height in inches")
	age := flag.Int("age", 0, "age in years")
	activity := flag.Int("activity", 0, "Activity Level Tier, where \n1 is Sedentary \n2 is Lightly Active \n3 is Moderately Active \n4 is Very Active \n5 is Extremely Active")

	flag.Parse()

	activityLevelModifier := convertActivityLevelTierToFunctionalValue(*activity)

	return float64(*bodyweight), float64(*height), float64(*age), activityLevelModifier
}

func convertActivityLevelTierToFunctionalValue(activityLevelTier int) float64 {
	var activityModifier float64
	if activityLevelTier == 1 {
		// 1. Sedentary (little to no exercise + work a desk job) = 1.2
		activityModifier = 1.2
	} else if activityLevelTier == 2 {
		// 2. Lightly Active (light exercise 1-3 days / week) = 1.375
		activityModifier = 1.55
	} else if activityLevelTier == 3 {
		// 3. Moderately Active (moderate exercise 3-5 days / week) = 1.55
		activityModifier = 1.55
	} else if activityLevelTier == 4 {
		// 4. Very Active (heavy exercise 6-7 days / week) = 1.725
		activityModifier = 1.725
	} else if activityLevelTier == 5 {
		// 5. Extremely Active (very heavy exercise, hard labor job, training 2x / day) = 1.9
		activityModifier = 1.9
	}
	return activityModifier
}

func main() {
	var bodyweight, heightInInches, age, activityModifier = convertFlagsToVars()

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

	printOutput(TDEE, proteinIntake, calFromProtein, carbIntake, calFromCarb, fatIntake, calFromFat)
}
