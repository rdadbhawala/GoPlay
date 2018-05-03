package main

func main() {
	// m := make(map[int]int)
	for ctr := 0; ctr < 5; ctr++ {
		analyze(ctr)
	}
}

func analyze(n int) {
	// nn := n * n
	grid := make([][]bool, n)
	for rowCtr := 0; rowCtr < n; rowCtr++ {
		grid[rowCtr] = make([]bool, n)
	}

}
