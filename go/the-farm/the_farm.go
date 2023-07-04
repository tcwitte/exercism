package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	totalAmount, err := fodderCalculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	factor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return totalAmount * factor / float64(cows), nil
}

var errInvalidNumberOfCows = errors.New("invalid number of cows")

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fodderCalculator, cows)
	} else {
		return 0, errInvalidNumberOfCows
	}
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (error *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", error.cows, error.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	} else {
		return nil
	}
}
