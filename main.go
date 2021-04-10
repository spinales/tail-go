package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	num     int
	numbers string // -n
	err     error
)

func main() {
	// flag.IntVar(&num, "n", 0, `This option shall be equivalent to -c number, except the starting location in the
	// file shall be measured in lines instead of bytes. The origin for counting shall be 1; that is, -n +1 represents
	// the first line of the file, -n -1 the last.

	// If neither -c nor -n is specified, -n 10 shall be assumed.`)
	flag.StringVar(&numbers, "n", "", `This option shall be equivalent to -c number, except the starting location in the 
	file shall be measured in lines instead of bytes. The origin for counting shall be 1; that is, -n +1 represents 
	the first line of the file, -n -1 the last.
	
	If neither -c nor -n is specified, -n 10 shall be assumed.`)
	flag.Parse()

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Fprint(os.Stderr, "No hay datos para recibir")
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
	switch numbers {
	case "":
		num = 10
	case "+1":
		fmt.Println(data[0])
		os.Exit(0)
	case "-1":
		fmt.Println(data[len(data)-1])
		os.Exit(0)
	default:
		num, err = strconv.Atoi(numbers)
		if err != nil {
			panic(err)
		}
	}

	switch {
	case num < len(data):
		fmt.Println(strings.Join(data[:len(data)-num], "\n"))
	default:
		fmt.Println(strings.Join(data, "\n"))
	}

	os.Exit(0)
}
