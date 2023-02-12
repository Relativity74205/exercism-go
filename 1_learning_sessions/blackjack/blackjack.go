package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
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
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Score := ParseCard(card1)
	card2Score := ParseCard(card2)
	playerScore := card1Score + card2Score
	dealerScore := ParseCard(dealerCard)

	switch {
	case card1Score == 11 && card2Score == 11:
		return "P"
	case playerScore == 21 && (dealerScore == 10 || dealerScore == 11):
		return "S"
	case playerScore == 21:
		return "W"
	case playerScore >= 17 && playerScore <= 20:
		return "S"
	case playerScore >= 12 && playerScore <= 16 && dealerScore >= 7:
		return "H"
	case playerScore >= 12 && playerScore <= 16 && dealerScore < 7:
		return "S"
	default:
		return "H"
	}
}
