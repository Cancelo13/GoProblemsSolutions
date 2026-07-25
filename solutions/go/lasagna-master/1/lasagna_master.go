package lasagnamaster

import "strings"

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	var noodles int = 0
	var sauce float64 = 0.0
	for i := 0 ; i < len(layers) ; i++ {
		if strings.ToLower(string(layers[i])) == "noodles"{
			noodles += 50
		}
		if strings.ToLower(string(layers[i])) == "sauce"{
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friend []string, mine []string) {
	mine[len(mine)-1] = string(friend[len(friend)-1])
}

func ScaleRecipe(receipe []float64, number int) []float64 {
	var newRec []float64 = []float64{}
	for i := 0; i < len(receipe); i++ {
		newRec = append(newRec, (receipe[i] * float64(number) / 2))
	}
	return newRec
}
