package leetcode

func maximumWealth(accounts [][]int) int {
	maxAccount, sumAccount := 0, 0

	for _, account := range accounts {

		for _, value := range account {
			sumAccount += value
		}

		if maxAccount < sumAccount {
			maxAccount = sumAccount
		}
		sumAccount = 0
	}

	return maxAccount
}
