package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	var nFlag = flag.Bool("n", false, "the trailing newline is suppressed")
	var color = flag.String("color", "white", "define the color of output")
	flag.Usage = usage
	flag.Parse()
	//flag.NFlag()
	for _, arg := range flag.Args() {
		if strings.HasPrefix("-", arg) {
			break
		}
		fmt.Fprintf(os.Stdout, replaceEnv(arg))
	}

	if *nFlag {
		fmt.Fprintf(os.Stdout, "\n")
	}

	if *color == "white" {
		fmt.Fprintf(os.Stdout, "white")
	}
}

func replaceEnv(s string) string {
	result := ""
	temp := ""
	for _, char := range s {
		if char != '$' {
			if temp == "" {
				result += string(char)
			} else if char == ' ' && char == '\t' && char == '\n' {
				tempResult := os.Getenv(temp[1:])
				if tempResult != "" {
					result += tempResult
				} else {
					result += temp
				}

				temp = ""

			} else {
				temp += string(char)
			}
		} else {
			temp = "$"
		}
	}

	if temp != "" {
		tempResult := os.Getenv(temp[1:])
		if tempResult != "" {
			result += tempResult
		} else {
			result += temp
		}
	}

	return result
}
