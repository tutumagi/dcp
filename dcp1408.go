package dcp

/**
Daily Coding Problem: Problem #1408
Given a list, sort it using this method: reverse(lst, i, j), which reverses lst from i to j.
*/

func reverse(lst []int, i, j int) []int {
	if i >= len(lst) || j >= len(lst) {
		return lst
	}
	for i < j {
		lst[i], lst[j] = lst[j], lst[i]
		i++
		j--
	}
	return lst
}
