package blackjack

var cards = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	if val, ok := cards[card]; ok {
		return val
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardsSum := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	switch {
	case cardsSum == 21:
		if dealerCardValue < 10 {
			return "W"
		}
		return "S"
	case cardsSum >= 17 && cardsSum <= 20:
		return "S"
	case cardsSum >= 12 && cardsSum <= 16:
		if dealerCardValue > 6 {
			return "H"
		}
		return "S"
	case cardsSum <= 11:
		return "H"
	default:
		return "P"
	}
}
