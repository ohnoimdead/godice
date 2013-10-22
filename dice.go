// Dice generates random numbers to simulate polyhedral dice.
//
// It takes input in the form of xdy where x is the quantity of dice and y is the number of sides.
// For example 4d6 will generate 4 random numbers between 1 and 6.
//
// Providing no input will re-roll the prior sequence.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	var dice string

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		print("Roll: ")
		fmt.Scanln(&dice)

		parts := strings.Split(dice, "d")
		if len(parts) != 2 {
			fmt.Println("Looking for something in the format of 4d6")
		} else {
			qty, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println(err)
			}
			sides, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println(err)
			}

			total := 0
			for i := 0; i < qty; i++ {
				val := r.Intn(sides) + 1

				if i > 0 {
					fmt.Print(" + ")
				}

				fmt.Print(strconv.Itoa(val))
				total += val
			}

			fmt.Println(" =", total)
		}
	}
}
