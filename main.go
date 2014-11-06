package main

import (
	"bufio"
	"fmt"
	"github.com/AnarchyLime/gcalc/gcalc"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input, result string
	var err error

	calc := gcalc.NewCalc()

	fmt.Println("Enter expression (enter q to quit):")

	for {
		input, _ = reader.ReadString('\n')

		if input[0:1] == "q" {
			fmt.Println("=> goodbye!")
			return
		}

		input = input + "="

		result, err = calc.ProcessExpr(input)
		if err != nil {
			fmt.Println("=> Error: " + err.Error())
		} else {
			fmt.Println("=> " + result)
		}
	}
}
