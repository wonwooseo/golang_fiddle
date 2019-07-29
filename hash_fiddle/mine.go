package main

import (
	"fmt"
	"crypto/sha256"
	"flag"
	"strings"
)

func main() {
	genesis := flag.String("i", "The quick brown fox jumps over the lazy dog", "Genesis input")
	target := flag.Int("r", 8, "Target difficulty")
	flag.Parse()
	mine(1, *target, *genesis)
}

func mine(difficulty int, target int, input string) {
	if difficulty > target {
		return
	}
	nonce := 1
	raw := sha256.Sum256([]byte(fmt.Sprintf("%d%s", nonce, input)))
	h := fmt.Sprintf("%x", raw)
	for h[0:difficulty] != strings.Repeat("0", difficulty) {
		nonce += 1
		raw = sha256.Sum256([]byte(fmt.Sprintf("%d%s", nonce, input)))
		h = fmt.Sprintf("%x", raw)
	}
	fmt.Printf("Found answer for difficulty %d\n", difficulty)
	fmt.Printf(" input=%s\n nonce=%d\n output=%s\n\n", input, nonce, h)
	mine(difficulty + 1, target, h)
}
