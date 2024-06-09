package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "one":
		return 1
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
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	//CardSum calculates the numerial sum of both player cards.
	cardSum := ParseCard(card1) + ParseCard(card2)

	var response string

	switch {
	case cardSum == 22:
		response = "P"
	case cardSum == 21:
		if ParseCard(dealerCard) != 11 && ParseCard(dealerCard) != 10 {
			response = "W"
		}
		response = "S"
	case cardSum >= 17 && cardSum <= 20:
		response = "S"
	case cardSum >= 12 && cardSum <= 16:
		if ParseCard(dealerCard) >= 7 {
			response = "H"
		} else {
			response = "S"
		}
	case cardSum <= 11:
		response = "H"
	}

	return response
}
