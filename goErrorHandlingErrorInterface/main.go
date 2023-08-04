package main

import "fmt"

type divideError struct {
	dividend    float64
	contextFunc string
}

func (d divideError) Error() (err string) {
	return fmt.Sprintf("You can not divid %v by 0", d.dividend)
}

func divide(dividend, divisor int) (Ans float64, err error) {
	if dividend == 0 {
		return 0, divideError{dividend: float64(dividend), contextFunc: "divide"}
	}
	return float64(dividend / dividend), nil
}

func main() {
	repeat := 'y'
	divisor := 0
	dividend := 0

	for repeat == 'y' {
		fmt.Print("Enter divisor: ")
		_, err := fmt.Scan(&divisor)
		if err != nil {
			fmt.Print(err)
		}

		fmt.Printf("Enter dividend: ")
		_, err = fmt.Scan(&dividend)
		if err != nil {
			fmt.Print(err)
		}

		ans, err := divide(dividend, divisor)
		if err != nil {
			fmt.Print("\n", err)
		}
		fmt.Printf("=> %.4f", ans)

		fmt.Print("\n########################################################")
		fmt.Print("\nDo you want to continue? Enter (y/n)")
		_, err = fmt.Scanf("%s", repeat)
		if err != nil {
			fmt.Printf("%s\n", err)
			repeat = 'n'
		}
	}

}
