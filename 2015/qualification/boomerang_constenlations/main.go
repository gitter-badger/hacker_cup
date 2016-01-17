package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ToInt(str string) int {
	n, _ := strconv.ParseInt(str, 10, 64)
	return int(n)
}

func Distance(x, y []int) int {
	run := x[0] - y[0]
	rise := x[1] - y[1]
	return run*run + rise*rise
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nights := ToInt(scanner.Text())

	for i := 0; i < nights; i++ {
		scanner.Scan()
		numberOfStars := ToInt(scanner.Text())
		constellations := [][]int{}
		for j := 0; j < numberOfStars; j++ {
			scanner.Scan()
			starsStr := strings.Split(scanner.Text(), " ")
			stars := []int{ToInt(starsStr[0]), ToInt(starsStr[1])}
			constellations = append(constellations, stars)
		}
		boomerangs := 0
		for x := 0; x < len(constellations); x++ {
			for y := 0; y < len(constellations); y++ {
				for z := 0; z < len(constellations); z++ {
					if x != y && x != z && y != z {
						if Distance(constellations[x], constellations[y]) == Distance(constellations[y], constellations[z]) {
							boomerangs = boomerangs + 1
						}

					}
				}
			}
		}
		fmt.Printf("Case #%d: %d\n", i+1, boomerangs/2)
	}
}
