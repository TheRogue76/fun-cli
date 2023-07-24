package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func rgb(i int) (int, int, int) {
	f := func(offset int) int {
		return int(math.Sin(0.1*float64(i)+(2.0*float64(offset))*math.Pi/3)*127 + 128)
	}
	return f(0), f(1), f(2)
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes")
		fmt.Println("Usage: fortune | gorainbow")
	}

	reader := bufio.NewReader(os.Stdin)
	j := 0
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, input)
		j++
	}
}
