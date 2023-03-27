package main

import "testing"

func TestDaily(t *testing.T) {
	a1 := countEven(5)
	t.Log(a1)

	a2 := minOperations([]int{1, 2, 3}, 2)
	t.Log(a2)

	a3 := digitCount("1312312131")
	t.Log(a3)

	a4 := rearrangeCharacters("sadadada", "Asdaxcscas")
	t.Log(a4)

	a5 := areSentencesSimilar("awdada", "dwqqdq")
	t.Log(a5)

	a6 := countNicePairs([]int{42, 11, 1, 97})
	t.Log(a6)

	a7 := waysToMakeFair([]int{2, 1, 6, 4})
	t.Log(a7)

	a8 := countAsterisks("l|*e*et|c**o|*de|")
	t.Log(a8)

	t.Log("test")
}
