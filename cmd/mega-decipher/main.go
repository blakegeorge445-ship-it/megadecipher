package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: mega-decipher <ciphered-url>")
		os.Exit(1)
	}

	url := os.Args[1]
	fmt.Println("Ciphered link =>", url)

	decipheredUrl, err := Decipher(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Original link =>", decipheredUrl)
}

func Decipher(url string) (string, error) {
	url = strings.TrimSpace(url)
	if url == "" {
		return "", fmt.Errorf("invalid Mega ciphered URL")
	}

	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "www.")

	prefixes := []string{
		"mega.nz/#!",
		"mega.co.nz/#!",
	}

	for _, prefix := range prefixes {
		if !strings.HasPrefix(url, prefix) {
			continue
		}

		remainder := strings.TrimPrefix(url, prefix)
		parts := strings.SplitN(remainder, "!", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return "", fmt.Errorf("invalid Mega ciphered URL")
		}

		return fmt.Sprintf("https://mega.nz/file/%s#%s", parts[0], parts[1]), nil
	}

	return "", fmt.Errorf("invalid Mega ciphered URL")
}
