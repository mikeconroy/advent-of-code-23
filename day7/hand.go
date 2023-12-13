package day7

type Hand struct {
	cards []int
	bid   int
	hType int
}

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

func (hand Hand) getType() int {
	if hand.cards[0] == hand.cards[1] &&
		hand.cards[0] == hand.cards[2] &&
		hand.cards[0] == hand.cards[3] &&
		hand.cards[0] == hand.cards[4] {
		return FIVE_OF_A_KIND
	}

	// Use a map [faceCard] -> Count
	// [2] = 2
	// Then use the map to work out the highest type.
	valCount := make(map[int]int)
	valsWithPairOrBetter := make(map[int]bool)
	for _, card := range hand.cards {
		valCount[card] = valCount[card] + 1
		if valCount[card] > 1 {
			valsWithPairOrBetter[card] = true
		}
	}

	highestType := HIGH_CARD
	for _, count := range valCount {

		if count == 5 {
			return FIVE_OF_A_KIND
		} else if count == 4 {
			highestType = FOUR_OF_A_KIND
		} else if count == 3 {
			if len(valsWithPairOrBetter) > 1 {
				if highestType < FULL_HOUSE {
					highestType = FULL_HOUSE
				}
			} else if highestType < THREE_OF_A_KIND {
				highestType = THREE_OF_A_KIND
			}
			// Check for full house.
		} else if count == 2 {
			// Check for two pair
			if len(valsWithPairOrBetter) > 1 {
				if highestType < TWO_PAIR {
					highestType = TWO_PAIR
				}
			} else if highestType < ONE_PAIR {
				highestType = ONE_PAIR
			}
		}
	}

	return highestType
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
