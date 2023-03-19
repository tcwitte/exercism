package purchase

import "fmt"

const CAR = "car"
const TRUCK = "truck"
const DEPRECIATION_LT3YR = 0.8
const DEPRECIATION_LT10YR = 0.7
const DEPRECIATION_10YR = 0.5

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == CAR || kind == TRUCK
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var choice string
	if option1 < option2 {
		choice = option1
	} else {
		choice = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", choice)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3.0 {
		return originalPrice * DEPRECIATION_LT3YR
	} else if age < 10.0 {
		return originalPrice * DEPRECIATION_LT10YR
	} else {
		return originalPrice * DEPRECIATION_10YR
	}
}
