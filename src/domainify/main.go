// main.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {

		text := s.Text()

		var newText []rune

		for _, ch := range text {
			if unicode.IsSpace(ch) {
				ch = '-'
				continue
			}

			if !strings.ContainsRune(allowedChars, ch) {
				continue
			}

			newText = append(newText, ch)
		}

		fmt.Println(string(newText))
	}
}
