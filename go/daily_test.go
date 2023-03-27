package main

import "testing"

func TestDaily(t *testing.T) {
	countEven(5)
	minOperations([]int{1, 2, 3}, 2)
	digitCount("1312312131")
	rearrangeCharacters("sadadada", "Asdaxcscas")
	areSentencesSimilar("awdada", "dwqqdq")
	countNicePairs([]int{42, 11, 1, 97})
	waysToMakeFair([]int{2, 1, 6, 4})
	countAsterisks("l|*e*et|c**o|*de|")
}
