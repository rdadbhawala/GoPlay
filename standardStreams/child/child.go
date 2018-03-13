package main

import (
	"io"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	len, _ := strconv.Atoi(os.Args[1])
	wg.Add(2)
	go dump(len, os.Stdout, "0123456789")
	go dump(len, os.Stderr, "ABCDEFGHIJ")
	wg.Wait()
}

func dump(len int, dst io.Writer, str string) {
	defer wg.Done()
	for ctr := 0; ctr < len; ctr += 10 {
		io.WriteString(dst, str)
	}
}
