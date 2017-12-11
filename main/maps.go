package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56,
		"charlie": 23, "delta": 87, "echo": 56,
		"foxtrot": 12, "golf": 34, "hotel": 16,
		"indio": 87, "juliet": 65, "kili": 43,
		"lima": 98}
)

func initMaps() {
	var map1 map[string]string
	map1 = make(map[string]string)
	map2 := map[string]string{"hello": "world", "seeyou": "tomorrow"}
	map3 := map[string]string{"aa": "aaa", "bb": "bbb"}
	map4 := map3

	map1["key1"] = "aa"
	map1["key2"] = "bb"
	map3 = map1
	delete(map2, "hello")
	if _, ok := map2["hello"]; ok {
		fmt.Println("map2[\"hello\"] = ", map2["hello"])
	} else {
		fmt.Println("map2[\"hello\"] not exists")
	}

	typeCheck(map1, map2, map3, map4)
}

//map4Funcs produces many func(int) int functions
func map4Funcs(n int) map[int]func(int) int {
	mf := make(map[int]func(int) int, n)
	for i := 1; i <= n; i++ {
		i2 := i
		//i increases in every loop
		//in every iteration, a closure makes reference to i
		//so after the loop ends, all the closures refers to the same i, which equals to n
		//so, assign i to another identifier i2
		//and every closure refers to a different value(different i2s)
		//then the closures are different
		mf[i] = func(x int) int { return i2 * x }
	}
	return mf
}

func mapSort(mss map[string]string) []string {
	if len(mss) == 0 {
		return make([]string, 0)
	}
	strings := make([]string, len(mss))
	i := 0
	for k := range mss {
		strings[i] = k
		i++
	}
	sort.Strings(strings)
	return strings
}

func mapInvKV(msi map[string]int) map[int][]string {

	miss := make(map[int][]string)
	for s, i := range msi {
		miss[i] = append(miss[i], s)
	}
	return miss
}
