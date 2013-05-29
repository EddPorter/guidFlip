package main

import (
	"fmt"
	"os"
	"regexp"
)

func flip(guid string) (output string) {
	var validGuid = regexp.MustCompile(`^([0-9a-fA-F]{8})-([0-9a-fA-F]{4})-([0-9a-fA-F]{4})-([0-9a-fA-F]{4})-([0-9a-fA-F]{12})$`)
	var guidMatch = validGuid.FindStringSubmatch(guid)

	var pairMatch = regexp.MustCompile(`..`)
	for n := 1; n <= 3; n++ {
		var pairs = pairMatch.FindAllString(guidMatch[n], -1)
		for i := len(pairs) - 1; i >= 0; i-- {
			output += pairs[i]
		}
	}
	output += guidMatch[4]
	output += guidMatch[5]
  return
}

func main() {
	var guid string = os.Args[1]
	fmt.Println(flip(guid))
}
