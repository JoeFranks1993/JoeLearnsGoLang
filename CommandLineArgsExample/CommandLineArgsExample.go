package main

import (
    "fmt"
	"os"
	"strconv"
)

func main() {
	// Assign command line args to local variable 
	var args = os.Args;

	for i := 0; i < len(args); i++ {
		var iteratorAsString = strconv.Itoa(i);
		fmt.Println("Param # " + iteratorAsString + " = " + args[i]);
	}
}