package techpalace

import(
	"strings"
)

func WelcomeMessage(name string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(name)
}

func AddBorder(message string, stars int) string {
	var border string = ""
	for i := 0; i < stars; i++ {
		border += "*"
	}
	return border + "\n" + message + "\n" + border
}

func CleanupMessage(message string) string {
	var newMessage string = ""
	for i := 0; i < len(message); i++ {
		if message[i] != '*' && message[i] != '\n' {
			newMessage += message[i:i+1]
		}
	}
	var start bool = false
	message = ""
	for i := 0; i < len(newMessage); i++ {
		if newMessage[i] != ' ' {
			message += newMessage[i:i+1]
			start = true
		}
		if newMessage[i] == ' ' && start {
			message += newMessage[i:i+1]
		}
	}
	var i int = len(message) - 1
	for ; i >= 0; i-- {
		if message[i] != ' ' {
			break
		}
	}
	return message[0:i+1]
}