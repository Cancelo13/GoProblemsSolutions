package blackjack

func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

func FirstTurn(cardA, cardB, cardC string) string {
	var valA int = ParseCard(cardA)
	var valB int = ParseCard(cardB)
	var valC int = ParseCard(cardC)
	var sum int = valA + valB
	switch {
	case cardA == "ace" && cardB == "ace":
		return "P"
	case sum == 21 && cardC != "ace" && valC != 10:
		return "W"
	case sum == 21 && (cardC == "ace" || valC == 10):
		return "S"
	case sum >= 17 && sum <= 20:
		return "S"
	case sum >= 12 && sum <= 16 && valC < 7:
		return "S"
	case sum >= 12 && sum <= 16 && valC >= 7:
		return "H"
	case sum <= 11:
		return "H"
	default:
		return ""
	}
}
