package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"strings"
)

func main() {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Aborted", err)
		return
	}
	genesis := flag.String("i", base64.StdEncoding.EncodeToString(b), "Genesis input")
	target := flag.Int("r", 6, "Target difficulty")
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
		nonce++
		raw = sha256.Sum256([]byte(fmt.Sprintf("%d%s", nonce, input)))
		h = fmt.Sprintf("%x", raw)
	}
	fmt.Printf("Found answer for difficulty %d\n", difficulty)
	fmt.Printf(" input=%s\n nonce=%d\n output=%s\n\n", input, nonce, h)
	mine(difficulty+1, target, h)
}
