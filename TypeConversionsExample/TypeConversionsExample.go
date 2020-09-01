package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	stringAsInteger := "64"
	stringToIntResult, err := stringToInt(stringAsInteger)

	myInteger := 50
	myIntegerConvertedToString := intToString(myInteger)

	boolAsString := "True"
	boolConvertedFromString, err := stringToBool(boolAsString)

	intWeWantToMakeAFloat := 500
	floatConvertedFromInt := intToFloat(intWeWantToMakeAFloat)

	floatWeWantToMakeAnInt := 50.7
	intConvertedFromFloat := floatToInt(floatWeWantToMakeAnInt)

	// print so the linter stops complaining the declared variables are not used, this is only for an example on converstion
	fmt.Println(stringToIntResult)
	fmt.Println(myIntegerConvertedToString)
	fmt.Println(boolConvertedFromString)
	fmt.Println(floatConvertedFromInt)
	fmt.Println(intConvertedFromFloat)
	fmt.Println(err)
}

func stringToInt(input string) (output int, err error) {

	output, err = strconv.Atoi(input) // Conversion from string to int has the potential to error, we need to capture the err value as well

	if err != nil {
		log.Fatal("Could not convert input string to integer. Input string = "+input, err)
	}

	return
}

func intToString(input int) (output string) {
	output = strconv.Itoa(input) // no error here because all int values will be guaranteed to be convertable to a string
	return
}

func stringToBool(input string) (output bool, err error) {

	output, err = strconv.ParseBool(input)

	if err != nil {
		log.Fatal("Could not convert input string to bool. Input string = "+input, err)
	}

	return
}

func intToFloat(input int) (output float64) {
	output = float64(input)
	return
}

func floatToInt(input float64) (output int64) {
	output = int64(input)
	return
}
