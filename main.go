package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

var matrix int = 32
var grid [32]*[]int

func main() {
	getGrid()
	for {
		clearScreen()

		var numTotal string

		fmt.Println("Randomio")
		fmt.Println("It is a random number generator that generates random numbers from a matrix of 1s and 0s and converts them to decimal.")
		fmt.Println("Matrix Shape: ", matrix)
		fmt.Println("\n\n\n ")

		for i := 0; i < matrix; i++ {
			updateRow(rand.Intn(matrix))
			numStr := ""
			for j := 0; j < matrix; j++ {
				fmt.Print((*grid[i])[j], "    ")
				numStr += fmt.Sprint(strconv.Itoa((*grid[i])[j]))
			}

			num, err := strconv.ParseInt(numStr, 2, 64)
			if err != nil {
				panic(err)
			}

			numTotal += strconv.Itoa(int(num))
			fmt.Print("value: ", num, "\n")
		}

		fmt.Print("\n\n\n\n\n\n\n")
		fmt.Printf("RESULT: %s \n\nLENGTH: %d\n", numTotal, len(numTotal))

		<-time.Tick(1 * time.Second)
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	data, _ := cmd.Output()
	fmt.Println(string(data))
}

func getGrid() {
	for i := 0; i < matrix; i++ {
		updateRow(i)
	}
}

func updateRow(num int) {

	lines := []int{}
	for j := 0; j < matrix; j++ {
		lines = append(lines, rand.Intn(2))
	}

	grid[num] = &lines

}
