package main

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