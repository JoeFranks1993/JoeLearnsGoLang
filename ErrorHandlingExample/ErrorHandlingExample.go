package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Get the param passed in via the command line.
	var commandLineInput = os.Args[1]

	// Parse the input to an int.
	var inputAsInt, inputConversionErr = strconv.Atoi(commandLineInput)

	if inputConversionErr == nil { // If the conversion from string to int failed, do nothing and print the error
		var result, err = DoAThing(inputAsInt) // Call our custom function
		if err == nil {                        // If we did not get an error, print the result out put to the console after casting it back to a string. Otherwise print the error
			fmt.Println("The input value plus 100 is " + strconv.Itoa(result))
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(inputConversionErr)
	}
}

func DoAThing(inputValue int) (result int, err error) {
	if inputValue < 10 {
		err = errors.New("Values less than 10 cannot be used")
		return
	} else {
		result = inputValue + 100
		return
	}
}
