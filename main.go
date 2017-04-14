package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/cznic/sortutil"
)

func main() {
	escape := flag.Bool("escape", false, "print escaped characters (Go syntax)")
	flag.Parse()

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
		if *escape {
			fmt.Printf("%q", ch)
		} else {
			fmt.Printf("%c", ch)
		}
	}

	fmt.Printf("\n")
}
