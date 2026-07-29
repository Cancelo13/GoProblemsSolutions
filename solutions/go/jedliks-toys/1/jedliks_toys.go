package jedlik

import "fmt"

func (car *Car) Drive(){
	if car.battery < car.batteryDrain{
		return
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%s", car.battery, "%")
}

func (car Car) CanFinish(trackDistance int) bool{
	var time float64 = float64(trackDistance) / float64(car.speed)
	return time*float64(car.batteryDrain) <= float64(car.battery)
}


