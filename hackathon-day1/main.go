package main

import "fmt"

func main() {

	for {
		var A, B int
		fmt.Print("Input A: ")
		a, err := fmt.Scanf("%v",&A)
		if err != nil {
			fmt.Println("=ENTER A VALID NUMBER=")
			continue
		}


		fmt.Print("Input B: ")
		c, err := fmt.Scanf("%v",&B)
		if err != nil {
			fmt.Println("=ENTER A VALID NUMBER=")
			continue
		}
		
		fmt.Println(a,c)

		var operation int
		fmt.Println("1: Addition")
		fmt.Println("2: Subtraction")
		fmt.Println("3: Multiplication ")
		fmt.Println("4: Division")
		fmt.Println("5: Help")
		fmt.Println("6: Quit")
		fmt.Print("Choice Operation: ")
		fmt.Scan(&operation)

		switch operation {
		case 1:

			fmt.Println("=ANSWER=: ", fmt.Sprint(A+B))
			continue

		case 2:

			fmt.Println("=ANSWER=: ", fmt.Sprint(A-B))
			continue

		case 3:

			fmt.Println("=ANSWER=: ", fmt.Sprint(A*B))
			continue

		case 4:
			if B == 0  {
				fmt.Println("=ERROR=")
			} else {
				fmt.Println("=ANSWER=:", A/B,"=REMAINDER=:", A%B)
			}
			continue
		case 5:
			fmt.Println("Welcome To Pro Claculator")
			fmt.Println("What do you want to claculate today?")
			fmt.Println("We help in addition, subtraction, multiplication and division.")
			continue

		case 6:
			fmt.Println("==THANK YOU FOR CLACULATING WITH US==-->")
			break

		default:
			fmt.Println("=OUT OF OPERATING RANGE=", "START AGAIN")
			continue
		}
		break
	}
}