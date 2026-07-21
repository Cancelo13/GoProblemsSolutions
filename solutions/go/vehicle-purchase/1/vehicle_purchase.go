package purchase

func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

func ChooseVehicle(modelA, modelB string) string {
	var model string 
	if modelA <= modelB {
		model = modelA
	}else{
		model = modelB
	}
	return model + " is clearly the better choice."
}

func CalculateResellPrice(price, age float64) float64 {
	if age < 3 {
		return price * 0.8
	} else if age >= 3 && age < 10 {
		return price * 0.7
	} else {
		return price * 0.5
	}
}
