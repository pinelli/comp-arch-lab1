package main

import "time"

var templ string

type byDistance []string

func (s byDistance) Len() int {
	return len(s)
}
func (s byDistance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byDistance) Less(i, j int) bool {
	return Distance(s[i], templ) < Distance(s[j], templ)
}

func fillColoumn(len_s1 int) []int {
	defer FuncTimeTrack(time.Now())

	column := make([]int, len_s1+1)
	for y := 1; y <= len_s1; y++ {
		column[y] = y
	}
	return column
}

func Y(x int, column []int, s1 []rune, s2 []rune, lastdiag *int) {
	defer FuncTimeTrack(time.Now())

	var cost, olddiag, len_s1 int
	len_s1 = len(s1)
	for y := 1; y <= len_s1; y++ {
		olddiag = column[y]
		cost = 0
		if s1[y-1] != s2[x-1] {
			cost = 1
		}
		column[y] = min(
			column[y]+1,
			column[y-1]+1,
			*lastdiag+cost)
		*lastdiag = olddiag
	}
}
func Distance(str1, str2 string) int {
	defer FuncTimeTrack(time.Now())

	var lastdiag int
	s1, s2 := []rune(str1), []rune(str2)
	len_s1, len_s2 := len(s1), len(s2)
	column := fillColoumn(len_s1)

	for x := 1; x <= len_s2; x++ {
		column[0] = x
		lastdiag = x - 1
		Y(x, column, s1, s2, &lastdiag)
	}
	return column[len_s1]
}

func min(a, b, c int) int {
	defer FuncTimeTrack(time.Now())

	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
