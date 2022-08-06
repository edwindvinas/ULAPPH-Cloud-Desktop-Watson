package main

import (
"os"
"fmt"
"io/ioutil"
//"bufio"
//"bytes"
"strings"
"github.com/tidwall/pretty"
)

func main() {
	arg := os.Args[1] //file
	arg2 := os.Args[2] //target skill
	
	fmt.Printf("ARG: %v\n", arg)
	fmt.Printf("ARG2: %v\n", arg2)
	fmt.Printf("File: %v\n", arg)
	fmt.Printf("\n++ Opening config file\n")
	watsfile, err := os.Open(arg)
	if err != nil {
		panic(err)
	} 
	fmt.Printf("\n++ Successfully Opened input file\n")
	defer watsfile.Close()
	byteInput, err := ioutil.ReadAll(watsfile)
	if err != nil {
		panic(err)
	}
	//var buf bytes.Buffer
	result := pretty.Pretty(byteInput)
	newFile := strings.Replace(arg, ".json", "", -1)
	err = ioutil.WriteFile(fmt.Sprintf("%v_(beautified).json", newFile), []byte(result), 0644)
    if err != nil {
		panic(err)
    }
}

