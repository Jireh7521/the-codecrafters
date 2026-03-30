package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("THE BASE CONVERTER")
	for {
		var hex, bin, dec string

		var bases int
		fmt.Println("1: HEX")
		fmt.Println("2: BIN")
		fmt.Println("3: DEC")
		fmt.Println("4: QUIT")
		fmt.Print("=CHOICE BASES=: ")
		fmt.Scan(&bases)

		switch bases {
		case 1:
			fmt.Print("=INPUT HEX VALUE=: ")
			fmt.Scan(&hex)

			dec, err := strconv.ParseInt(hex, 16, 64)
			if err != nil {
				fmt.Println("=INVALID HEX VALUE=")
			}
			fmt.Println("=ANSWER=", dec)
			continue

		case 2:
			fmt.Print("=INPUT BIN VALUE=: ")
			fmt.Scan(&bin)

			dec, err := strconv.ParseInt(bin, 2, 64)
			if err != nil {
				fmt.Println("=INVALID BIN VALUE=")
			}
			fmt.Println("=ANSWER=", dec)
			continue

		case 3:
			fmt.Print("=INPUT DEC VALUE=: ")
			fmt.Scan(&dec)
			
			dec, err := strconv.ParseInt(dec, 10, 64)
			if err != nil {
				fmt.Println("=INVALID DEC VALUE=")
			}
			hexN := strconv.FormatInt(dec, 16)
			binN := strconv.FormatInt(dec, 2)
			fmt.Println("=ANSWER=", binN)
			fmt.Println("=ANSWER=", strings.ToUpper(hexN))

			continue

		case 4:
			fmt.Println("=THANK YOU FOR YOUR TIME=-->")
			break

		default:
			fmt.Println("=*OUT OF BASES RAGNE", "START AGAIN*=")
			continue
		}
		break
	}

}