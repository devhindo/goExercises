package main

/*

sort string in ascending order 

*/

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	t := strings.Split(s, "")
	sort.Strings(t)
	return strings.Join(t, "")
}

func main() {
	s := "devhindo"
	println(sortString(s))
}