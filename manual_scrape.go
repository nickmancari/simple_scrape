package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"bufio"
)

func getSite(s string) {
	response, err := http.Get(s)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		bodyText, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", bodyText)
	}
}

func main() {
	
	var reader = bufio.NewReader(os.Stdin)
	Input, _ := reader.ReadString()

	getSite(Input)
}
