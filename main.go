package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Error reading clipboard:", err)
		return
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-h", "--help":
			printHelp()
			return
		}
		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "-h", "--help":
				printHelp()
				return
			case "-p", "--print":
				text, err = createNewToken(text, os.Args)
				if err != nil {
					fmt.Println("Error creating token:", err)
					return
				}
				fmt.Println(text)
				return
			case "-v", "--verbose":
				fmt.Printf("Original text:\n%s\n\n", text)
				text, err = createNewToken(text, os.Args)
				if err != nil {
					fmt.Println("Error creating token:", err)
					return
				}
				fmt.Printf("New text:\n%s\n\n", text)
			}
		}
	}
	text, err = createNewToken(text, os.Args)
	if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}
	err = clipboard.WriteAll(text)
	if err != nil {
		fmt.Println("Error writing to clipboard:", err)
		return
	}
	fmt.Println("New string saved to clipboard!")
}

func getDoubleSidedTokens() map[string]string {
	doubleSidedTokens := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
		"<": ">",
		">": "<",
		"]": "[",
		"}": "{",
		")": "(",
	}
	return doubleSidedTokens
}

func matchDoubleSidedToken(token string) string {
	doubleSidedTokens := getDoubleSidedTokens()
	result, ok := doubleSidedTokens[token]
	if ok {
		return result
	}
	return token
}

func createNewToken(original string, args []string) (string, error) {
	doubleSidedTokens := getDoubleSidedTokens()
	result := ""
	switch {
	case len(args) > 1:
		token := args[1]
		token2 := token
		if token[0] == '-' {
			printHelp()
			err := fmt.Errorf("Improper order of arguments.\n")
			return original, err
		}
		_, found := doubleSidedTokens[token]
		if found {
			token2 = matchDoubleSidedToken(token)
		}
		original = strings.Trim(original, " \t\n")
		result = fmt.Sprintf("%s%s%s", token, original, token2)
	default:
		printHelp()
	}
	return result, nil
}

func printHelp() {
	fmt.Println("Usage:\nblsurround <surround token>")
	fmt.Printf("\nFlags:\n-h or --help : Show this menu\n-p or --print : print the result to the console instead of saving it back to the clipboard.\n\n")
}