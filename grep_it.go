package main

import (

	"bufio"
	"fmt"
	"os"
	"regexp"

)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)

	r, _ := regexp.Compile(os.Args[1])
	for scanner.Scan() { 
		if r.MatchString(scanner.Text()) {
			fmt.Println(scanner.Text())
		}
	}
}	
