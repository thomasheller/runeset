package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/cznic/sortutil"
)

func main() {
	runes := make(map[rune]bool)

	r := bufio.NewReader(os.Stdin)

	for {
		if ch, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			runes[ch] = true
		}
	}

	var set sortutil.RuneSlice

	for ch := range runes {
		set = append(set, ch)
	}

	sort.Sort(set)

	for _, ch := range set {
		fmt.Printf("%c", ch)
	}

	fmt.Printf("\n")
}
