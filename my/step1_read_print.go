package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("user> ")
	for scanner.Scan() {
		line := scanner.Text()
		result := rep(line)
		fmt.Println(result)
		fmt.Printf("user> ")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("could not read line from STDIN: %v", err)
	}

	fmt.Println("")
}

func read(str string) malType {
	return readStr(str)
}

func eval(mt malType) malType {
	return mt
}

func prnt(mt malType) string {
	return prStr(mt)
}

func rep(str string) string {
	readResult := read(str)
	evalResult := eval(readResult)
	prntResult := prnt(evalResult)

	return prntResult
}
