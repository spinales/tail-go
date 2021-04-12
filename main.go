package main

import (
	"bufio"
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
	f       string // -f
	err     error
)

func main() {
	flag.StringVar(&f, "f", "", `If the input file is a regular file or if the file operand specifies a FIFO, do not 
	terminate after the last line of the input file has been copied, but read and copy further bytes from the input 
	file when they become available. If no file operand is specified and standard input is a pipe, the -f option 
	shall be ignored. If the input file is not a FIFO, pipe, or regular file, it is unspecified whether or 
	not the -f option shall be ignored.`)
	flag.StringVar(&numbers, "n", "", `This option shall be equivalent to -c number, except the starting location in the 
	file shall be measured in lines instead of bytes. The origin for counting shall be 1; that is, -n +1 represents 
	the first line of the file, -n -1 the last.
	
	If neither -c nor -n is specified, -n 10 shall be assumed.`)
	flag.Parse()

	if f != "" {
		arr := fifo(f)
		switch {
		case num < len(arr):
			fmt.Println(strings.Join(arr[:len(arr)-num], "\n"))
		default:
			fmt.Println(strings.Join(arr, "\n"))
		}
		os.Exit(0)
	}

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

func fifo(file string) []string {
	var data []string

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	s, e := Readln(r)
	for e == nil {
		data = append(data, s)
		s, e = Readln(r)
	}

	return data
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
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
