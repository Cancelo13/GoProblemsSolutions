package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	}
	if balance < 1000 {
		return 0.5
	}
	if balance < 5000 {
		return 1.621
	}
	return 2.475
}

func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

func YearsBeforeDesiredBalance(balance float64, target float64) int {
	var cnt int = 0
	for balance < target {
		balance = AnnualBalanceUpdate(balance)
		cnt++
	}
	return cnt
}

