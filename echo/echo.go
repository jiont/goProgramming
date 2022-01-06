package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	nFlag:= flag.Bool("n", false, "the trailing newline is suppressed")
	//hFlag:= flag.Bool("h",false,"-n trailing newline is suppressed")
	flag.Parse()
	for index, arg := range os.Args[1:] {
		fmt.Fprintf(os.Stdout, arg)
		if index != len(os.Args)-2 {
			fmt.Fprintf(os.Stdout, " ")
			fmt.Printf("\033[1;31;40m%s\033[0m\n","Red.")
		}

	}
	if *nFlag{
		fmt.Fprintf(os.Stdout, "\n")
	}
}
