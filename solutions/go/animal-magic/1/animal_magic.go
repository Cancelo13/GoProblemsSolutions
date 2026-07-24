package chance

import "math/rand"

func RollADie() int {
	var n int = rand.Intn(21)
	if n == 0 {
		n = 1
	}
	return n
}

func GenerateWandEnergy() float64 {
	return rand.Float64() * 12.0
}

func ShuffleAnimals() []string {
	var animals []string = []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}