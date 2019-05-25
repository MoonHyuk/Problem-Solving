package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

	var mountainsCount int
	fmt.Fscanln(r, &mountainsCount)

	newHeight, maxHeight, kills, maxKills := 0, 0, 0, 0
	for i:= 0; i < mountainsCount; i ++ {
		fmt.Fscan(r, &newHeight)

		if newHeight > maxHeight {
			maxHeight = newHeight
			maxKills = max(maxKills, kills)
			kills = 0
		} else {
			kills ++
		}
	}
	maxKills = max(maxKills, kills)

	fmt.Fprintln(w, maxKills)
	w.Flush()
}
