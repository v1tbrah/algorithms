package design_an_atm_machine

// Task: https://leetcode.com/problems/design-an-atm-machine/

// Solution: https://leetcode.com/problems/design-an-atm-machine/submissions/1066032543/

// tags:
// #Array
// #Design

type ATM struct {
	idxToBanknotSize [5]int
	banknotesCount   [5]int
}

func Constructor() ATM {
	return ATM{
		idxToBanknotSize: [5]int{20, 50, 100, 200, 500},
		banknotesCount:   [5]int{},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, bc := range banknotesCount {
		this.banknotesCount[i] += bc
	}
}

// Time: O(n)
func (this *ATM) Withdraw(amount int) []int {
	tmp := this.banknotesCount

	used := make([]int, 5)
	for i := 4; i >= 0; i-- {
		bcSize := this.idxToBanknotSize[i]
		haveBanknotes := this.banknotesCount[i]
		// banknote < amount and have this banknotes
		if bcSize <= amount && haveBanknotes > 0 {
			countToUse := amount / bcSize
			if haveBanknotes < countToUse {
				countToUse = haveBanknotes
			}

			tmp[i] -= countToUse

			amount -= countToUse * bcSize

			used[i] = countToUse
		}
	}

	if amount != 0 {
		return []int{-1}
	}

	this.banknotesCount = tmp

	return used
}
