package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("plan.txt")
	defer file.Close()
	if err != nil {
		return
	}
	file.WriteString("I’m not afraid of difficulties and insist on learning programming")
	file1, _ := os.Open("plan.txt")
	defer file1.Close()
	buf := make([]byte, 1024)
	file1.Read(buf)
	fmt.Println(string(buf))
}
