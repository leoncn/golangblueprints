package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + " app",
	otherWord + " site",
	"get " + otherWord,
	"go " + otherWord,
	"lets " + otherWord,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		input := transforms[rand.Intn(len(transforms))]
		output := strings.Replace(input, otherWord, s.Text(), -1)

		fmt.Println(output)
	}
}
