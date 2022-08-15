package assigncookies

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	var children, cookieIndex int
	for {
		if children >= len(g) || cookieIndex >= len(s) {
			break
		}
		if g[children] <= s[cookieIndex] {
			children++
		}
		cookieIndex++
	}
	return children
}
