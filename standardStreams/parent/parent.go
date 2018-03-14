package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	outStr := &bytes.Buffer{}
	errStr := &bytes.Buffer{}

	cmd := exec.Command("child.exe", os.Args[1])
	cmd.Stderr = errStr
	cmd.Stdout = outStr
	cmd.Start()

	fmt.Println("out len", outStr.Len())
	fmt.Println("err len", errStr.Len())

	cmd.Wait()

	fmt.Println("out len", outStr.Len())
	fmt.Println(string(outStr.Bytes()))
	fmt.Println("err len", errStr.Len())
	fmt.Println(string(errStr.Bytes()))
}
