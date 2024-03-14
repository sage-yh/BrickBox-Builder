package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	width := 10
	height := 5

	box := make([][]rune, height)
	for i := range box {
		box[i] = make([]rune, width)
	}

	for i := range box {
		for j := range box[i] {
			box[i][j] = ' '
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printBox(box)
		fmt.Println("Enter coordinates (row column) to add a brick (e.g., 1 2), or 'q' to quit:")
		scanner.Scan()
		input := scanner.Text()
		if input == "q" {
			break
		}

		var row, col int
		_, err := fmt.Sscanf(input, "%d %d", &row, &col)
		if err != nil || row < 0 || row >= height || col < 0 || col >= width {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		box[row][col] = '#'
	}

	fmt.Println("Final Box:")
	printBox(box)
}

func printBox(box [][]rune) {
	for _, row := range box {
		for _, val := range row {
			fmt.Printf("%c", val)
		}
		fmt.Println()
	}
}
