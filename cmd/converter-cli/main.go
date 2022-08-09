package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(0)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatal("failed to read input")
	}

	inputF := scanner.Text()

	f, err := strconv.ParseFloat(inputF, 64)
	if err != nil {
		log.Fatal(err)
	}

	c := (f - 32) * 5 / 9
	fmt.Print(c)
}
