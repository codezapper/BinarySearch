package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("test_data.csv")

	b1 := make([]byte, 7)
	f.Seek(8, 0)
	n1, _ := f.Read(b1)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}
