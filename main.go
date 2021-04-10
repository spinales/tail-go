package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const defalt int = 10

var num int

func main() {
	// flag.IntVar(&num, "n", 0, "")
	// flag.Parse()

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Fprint(os.Stderr, "No ha datos para recibir")
		os.Exit(1)
	}
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(bytes), "\n")

	print(arr)
}

func print(data []string) {
	if len(data) < 10 {
		fmt.Println(strings.Join(data, "\n"))
		os.Exit(0)
	}

	// print default len
	fmt.Println(strings.Join(data[:len(data)-defalt], "\n"))
}
