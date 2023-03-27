package main

/*
eg1:
initialEnergy =
1
initialExperience =
1
energy =
[1,1,1,1]
experience =
[1,1,1,50]

See https://leetcode.cn/problems/minimum-hours-of-training-to-win-a-competition/ for more details
*/

/*
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	var hours int
	for i := 0; i < len(energy); i++ {
		if initialEnergy <= energy[i] {
			hours += energy[i] - initialEnergy + 1
			initialEnergy = 1
		} else {
			initialEnergy -= energy[i]
		}

		if initialExperience <= experience[i] {
			hours += experience[i] - initialExperience + 1
			initialExperience += experience[i] - initialExperience + 1
		}
		initialExperience += experience[i]
	}
	return hours
}
*/
