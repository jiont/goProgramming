package main


import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 :%v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	fmt.Printf("%d\n", len(counts))
}

func countLines(f *os.File, counts map[string]string, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		lastFileName, ok := counts[input.Text()]
		if ok {
			fmt.Printf("the line of %s of %s is the same as %s\n", input.Text(), fileName, lastFileName)
		}

		counts[input.Text()] = fileName
	}
}
