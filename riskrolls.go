package main

import (
	"fmt"
	"sort"
)

func sortRoll(r []int) {
	sort.Slice(r, func(i, j int) bool {
		return r[i] > r[j]
	})
}

func score(attacker, defender []int) (attackerLoss, defenderLoss int) {
	sortRoll(attacker)
	sortRoll(defender)
	for i := range defender {
		if attacker[i] > defender[i] {
			defenderLoss++
		} else {
			attackerLoss++
		}
	}
	return
}

func main() {
	diceFaces := []int{1, 2, 3, 4, 5, 6}
	tally := make(map[string]int)
	count := 0.0
	for _, a1 := range diceFaces {
		for _, a2 := range diceFaces {
			for _, a3 := range diceFaces {
				for _, d1 := range diceFaces {
					for _, d2 := range diceFaces {
						for _, d3 := range diceFaces {
							al, dl := score([]int{a1, a2, a3}, []int{d1, d2, d3})
							k := fmt.Sprintf("-%d -%d", al, dl)
							tally[k]++
							count++
						}
					}
				}
			}
		}
	}
	for k, t := range tally {
		fmt.Printf("%s %d %5.2f%%\n", k, t, float64(100.0*t)/count)
	}
}
