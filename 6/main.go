package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "hello friends"

	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.Contains(greetings, "helo"))
	fmt.Println(strings.ReplaceAll(greetings, "hello", "hi"))

	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "ll"))

	fmt.Println(strings.Split(greetings, " "))

	// the original value is unchanged
	fmt.Println("the original values is = ", greetings)

	// sort
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages) // it modify the existing slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
}
