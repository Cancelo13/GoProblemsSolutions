package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

func CanSpy(knightIsAwake bool, archerIsAwake bool, prisonerIsAwake bool) bool {
	return prisonerIsAwake || knightIsAwake || archerIsAwake
}

func CanSignalPrisoner(archerIsAwake bool, prisonerIsAwake bool) bool {
	return prisonerIsAwake && !archerIsAwake
}

func CanFreePrisoner(knightIsAwake bool, archerIsAwake bool, prisonerIsAwake bool, petDogIsPresent bool) bool {
	if (petDogIsPresent && !archerIsAwake){
		return true
	}
	if(!petDogIsPresent && prisonerIsAwake && !knightIsAwake && !archerIsAwake){
		return true
	}
	return false
}

