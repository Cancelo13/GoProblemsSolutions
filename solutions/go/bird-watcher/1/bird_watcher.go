package birdwatcher

func TotalBirdCount(birds []int)int{
	var sum int = 0
	for i := 0 ; i < len(birds) ; i++{
		sum += birds[i]
	}
	return sum
}

func BirdsInWeek(birds []int, week int) int {
	var start int = (week - 1) * 7
	var end int = week * 7
	var sum int = 0
	for i:= start ; i < end ; i++{
		sum += birds[i]
	}
	return sum
}

func FixBirdCountLog(birds []int) []int {
	for i := 0; i < len(birds) ; i+=2{
		birds[i] += 1
	}
	return birds
}
