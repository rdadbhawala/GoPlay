package main

import "fmt"

func main() {
	limit := 5
	start := []coordinate{coordinate{0, 0}}
	for ctr := 0; ctr < limit; ctr++ {
		start = calculateFor(start)
		newLen := len(start)
		fmt.Println(ctr+1, newLen, (ctr+1)*4+1)
		draw(ctr, start)
	}
}

func draw(n int, coors []coordinate) {
	halfSize := (n + 1) * 2
	size := (halfSize * 2) + 1
	grid := make([][]bool, size)
	for rowCtr := 0; rowCtr < size; rowCtr++ {
		grid[rowCtr] = make([]bool, size)
	}
	for i := range coors {
		coor := coors[i]
		grid[coor.x+halfSize][coor.y+halfSize] = true
	}
	for j := range grid {
		for k := range grid[j] {
			if grid[j][k] {
				fmt.Print("|")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func calculateFor(start []coordinate) []coordinate {
	result := []coordinate{}
	for startCtr := 0; startCtr < len(start); startCtr++ {
		currCoor := start[startCtr]
		x := currCoor.x
		y := currCoor.y
		result = insertInto(x+2, y+1, result)
		result = insertInto(x+2, y-1, result)
		result = insertInto(x-2, y+1, result)
		result = insertInto(x-2, y-1, result)
		result = insertInto(x+1, y+2, result)
		result = insertInto(x-1, y+2, result)
		result = insertInto(x+1, y-2, result)
		result = insertInto(x-1, y-2, result)
	}
	return result
}

func insertInto(newx, newy int, result []coordinate) []coordinate {
	ctrLen := len(result)
	matchFound := false
	for ctr := 0; ctr < ctrLen; ctr++ {
		// if same x,y .. don't insert
		currCoor := result[ctr]
		if currCoor.x == newx && currCoor.y == newy {
			matchFound = true
			break
		}
	}
	if !matchFound {
		result = append(result, coordinate{newx, newy})
	}
	return result
}

type coordinate struct {
	x, y int
}
