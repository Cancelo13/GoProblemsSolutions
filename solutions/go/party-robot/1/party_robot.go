package partyrobot

import "fmt"

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
}

func AssignTable(name string, table int, mateName string, direction string, distance float64) string {
	Welcome(name)
	var message string = fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	return Welcome(name) + "\n" + message + "\n" + fmt.Sprintf("You will be sitting next to %s.", mateName)
}
