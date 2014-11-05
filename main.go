package main

import (
	"bufio"
	"fmt"
	"github.com/AnarchyLime/gcalc/gcalc"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input, result string
	var err error
	r, _ := regexp.Compile(`^[\s\d\+-=]+$`)

	calc := gcalc.NewCalc()

	fmt.Println("Enter expression (enter q to quit):")

	for {
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "q" || input == "Q" {
			fmt.Println("=> goodbye!")
			return
		}

		input = input + "="

		if r.MatchString(input) {
			fields := strings.Fields(input)
			for _, f := range fields {
				for _, c := range f {
					result, err = calc.PushKey(string(c))
					if err != nil {
						break
					}
				}
				if err != nil {
					break
				}
			}
			if err != nil {
				fmt.Println("=> Error: " + err.Error())
			} else {
				fmt.Println("=> " + result)
			}
		} else {
			fmt.Println("=> Invalid Input: " + input)
		}
	}
}
