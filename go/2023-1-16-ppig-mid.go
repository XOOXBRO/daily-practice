package main

import (
	"strings"
)

// See https://leetcode.cn/problems/sentence-similarity-iii/ for more details
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1 := strings.Split(sentence1, " ")
	s2 := strings.Split(sentence2, " ")

	i, n1 := 0, len(s1)
	j, n2 := 0, len(s2)
	for i < n1 && i < n2 && s1[i] == s2[i] {
		i++
	}
	for j < n1-i && j < n2-i && s1[n1-j-1] == s2[n2-j-1] {
		j++
	}
	if i+j == n1 || i+j == n2 {
		return true
	}
	return false
}
