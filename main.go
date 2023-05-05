package main

import (
	"fmt"

	"github.com/marcelrieger/geometrie/punkt"
)

func main() {
	pt1 := punkt.Punkt{X: 2, Y: 3}
	fmt.Println(pt1)
	fmt.Println(pt1.Laenge())
	pt2 := punkt.Punkt{X: 3, Y: 4}
	fmt.Println(pt2)
	fmt.Println(pt1.Laenge())
}
