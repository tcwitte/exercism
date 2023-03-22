package blackjack

const ace = "ace"
const two = "two"
const three = "three"
const four = "four"
const five = "five"
const six = "six"
const seven = "seven"
const eight = "eight"
const nine = "nine"
const ten = "ten"
const jack = "jack"
const queen = "queen"
const king = "king"
const stand = "S"
const hit = "H"
const split = "P"
const automaticallyWin = "W"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case ace:
		return 11
	case two:
		return 2
	case three:
		return 3
	case four:
		return 4
	case five:
		return 5
	case six:
		return 6
	case seven:
		return 7
	case eight:
		return 8
	case nine:
		return 9
	case ten, jack, queen, king:
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardsSum := ParseCard(card1) + ParseCard(card2)

	switch {
	case card1 == ace && card2 == ace:
		return split
	case cardsSum == 21 && ParseCard(dealerCard) < 10:
		return automaticallyWin
	case cardsSum == 21 && ParseCard(dealerCard) >= 10:
		return stand
	case cardsSum >= 17 && cardsSum <= 20:
		return stand
	case cardsSum >= 12 && cardsSum <= 16 && ParseCard(dealerCard) >= 7:
		return hit
	case cardsSum >= 12 && cardsSum <= 16 && ParseCard(dealerCard) < 7:
		return stand
	default:
		return hit
	}
}
