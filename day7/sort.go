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
