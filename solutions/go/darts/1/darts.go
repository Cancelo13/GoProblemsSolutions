package darts

func Score(x, y float64) int{
	var distance float64 = x * x + y * y // Distance from the Center (0,0) => (x - 0)^2 + (y - 0)^2 = r ^ 2
	var score int = 0
	if distance <= 1 {
		score = 10
	}else if distance <= 25 {
		score = 5
	}else if distance <= 100{
		score = 1
	}
	return score
}