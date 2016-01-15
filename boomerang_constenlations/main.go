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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    nights := ToInt(scanner.Text())

    for i := 0; i < nights; i++ {
        scanner.Scan()
        numberOfStars := ToInt(scanner.Text())
        for j := 0; j < numberOfStars; j++ {
            scanner.Scan()
            starsStr := strings.Split(scanner.Text(), " ")
            stars := []int{ToInt(starsStr[0]), ToInt(starsStr[1])}
            fmt.Println(stars)
        }
    }
}
