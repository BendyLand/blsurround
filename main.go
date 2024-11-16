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
	err = handleArgs(os.Args, text)
	if err != nil {
		switch err.Error() {
		case "-h", "-p":
			{
			} // do nothing
		default:
			printUsage()
			fmt.Printf("%s\n\n", err)
		}
		return
	}
	text, err = createNewToken(text, os.Args)
	if err != nil {
		fmt.Printf("Error creating token: %s\n", err)
		printUsage()
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
		printUsage()
	}
	return result, nil
}

func printUsage() {
	fmt.Println("Usage:\nblsurround <surround token>")
	fmt.Printf(
		"\nFlags:" +
			"\n-h or --help    : Show this menu." +
			"\n-p or --print   : Print the result to the console instead of saving it back to the clipboard." +
			"\n-v or --verbose : Print the original text, print the new text, and save the new text to the clipboard.\n\n",
	)
}

func printHelpMenu() {
	fmt.Println("Bland Surround Tool will read the system clipboard and apply the given token on either side.")
	fmt.Println("\nFor example, if 'test' is saved to the system clipboard, running:\n`blsurround \\(`\nWill produce:\n'(test)'")
	fmt.Println("\nUsage:\nblsurround <surround token>")
	fmt.Printf(
		"\nFlags:" +
			"\n-h or --help    : Show this menu." +
			"\n-p or --print   : Print the result to the console instead of saving it back to the clipboard." +
			"\n-v or --verbose : Print the original text, print the new text, and save the new text to the clipboard.\n\n",
	)

}

func handleArgs(args []string, text string) error {
	if len(args) > 1 {
		switch args[1] {
		case "-h", "--help":
			printHelpMenu()
			err := fmt.Errorf("-h")
			return err
		}
		if len(args) > 2 {
			switch args[2] {
			case "-h", "--help":
				printHelpMenu()
				err := fmt.Errorf("-h")
				return err
			case "-p", "--print":
				text, err := createNewToken(text, os.Args)
				if err != nil {
					message := fmt.Errorf("Error creating token: %s\n", err)
					return message
				}
				fmt.Println(text)
				message := fmt.Errorf("-p")
				return message
			case "-v", "--verbose":
				fmt.Printf("Original text:\n%s\n\n", text)
				text, err := createNewToken(text, os.Args)
				if err != nil {
					message := fmt.Errorf("Error creating token: %s\n", err)
					return message
				}
				fmt.Printf("New text:\n%s\n\n", text)
			}
		}
	}
	return nil
}
