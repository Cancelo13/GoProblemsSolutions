package thefarm

import (
    "errors"
    "fmt"
)

func DivideFood(fodder FodderCalculator, cows int) (float64, error) {
	tot, err := fodder.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	fac, err2 := fodder.FatteningFactor()
	if err2 != nil {
		return 0, err2
	}
	return (tot / float64(cows)) * fac, nil
}

func ValidateInputAndDivideFood(fodder FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fodder, cows)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct{
	cows int
	message string
}

func (e *InvalidCowsError) Error() string {
  return fmt.Sprintf("%v cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error{
	if cows < 0 {
		return &InvalidCowsError{
			cows: cows,
			message: "there are no negative cows",
		}
	}
	if cows == 0 {
		return &InvalidCowsError{
			cows: cows,
			message: "no cows don't need food",
		}
	}
	return nil
}
