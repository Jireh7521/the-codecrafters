package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"regexp"
)

func ToUpper(text string) string {
	return strings.ToUpper(text)
}

func ToLower(text string) string {
	return strings.ToLower(text)
}

func Capitalise(text string) string {
	u := strings.ToLower(text)
	return strings.Title(u)
}

func Title(text string) string {
    return strings.Title(text)
}

func Snake(text string) string {
    text = strings.ToLower(text)
    re := regexp.MustCompile(`[^a-z0-9\s]`)
    text = re.ReplaceAllString(text, "")
    return strings.Join(strings.Fields(text), "_")
}


func Reverse(text string) string {
    words := strings.Fields(text)

    return strings.Join(words, " ")
}


func main() {
	
	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
	var text string

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
		

		var choice int
		fmt.Println("1: ToUpper")
		fmt.Println("2: ToLower")
		fmt.Println("3: Capitalise")
		fmt.Println("4: Title")
		fmt.Println("5 Snake")
		fmt.Println("6: Reverse")
		fmt.Println("7: Exit")
		fmt.Print("=CHOiCE NUMBER=:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if text == "" {
				fmt.Println("✗ No text provided. Usage: upper <text>")
			} else {
				fmt.Println("→", ToUpper(text))
			}
			continue

		case 2:
			if text == "" {
				fmt.Println("✗ No text provided. Usage: lower <text>")
			} else {
			fmt.Println("→", ToLower(text))
			}
			continue

		case 3:
			if text == "" {
				fmt.Println("✗ No text provided. Usage: cap <text>")
			} else {
				fmt.Println("→", Capitalise(text))
			}
			continue
			
		case 4:
			if text == "" {
				fmt.Println("✗ No text provided. Usage: title <text>")
			} else {
				fmt.Println("→", Title(text))
			}
			continue

		case 5:
			if text == "" {
				fmt.Println("✗ No text provided. Usage: snake <text>")
			} else {
				fmt.Println("→", Snake(text))
			}
			continue

		case 6:
			if text == "" {
				fmt.Println("✗ No text provided. Usage: reverse <text>")
			} else {
				fmt.Println("→", Reverse(text))
			}
			continue

		case 7:
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
			

		default:
			fmt.Printf("✗ Unknown command: %q\n", reader)
            fmt.Println("  Valid commands: upper, lower, cap, title, snake, reverse, exit")

			continue
		}

		break
		
	}
}
