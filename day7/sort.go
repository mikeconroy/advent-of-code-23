package day7

func sort(hands []Hand) []Hand {
	n := len(hands)
	swapped := true

	// Loop until there are no more swaps made (sorted).
	for swapped {
		swapped = false
		for i := 1; i <= n-1; i++ {
			if hands[i-1].strongerThan(hands[i]) {
				swapHand := hands[i-1]
				hands[i-1] = hands[i]
				hands[i] = swapHand
				swapped = true
			}
		}
	}
	return hands
}

// Returns true if hand1 is stronger than hand2.
func (h1 Hand) strongerThan(h2 Hand) bool {
	h1Type := h1.hType
	h2Type := h2.hType
	if h1Type > h2Type {
		return true
	} else if h1Type < h2Type {
		return false
	} else {
		// Both are equal types so compare card face values in order.
		for i := 0; i < len(h1.cards); i++ {
			if h1.cards[i] > h2.cards[i] {
				return true
			} else if h1.cards[i] < h2.cards[i] {
				return false
			}
		}
	}
	// fmt.Println(h1Type, h2Type)
	return false
}
