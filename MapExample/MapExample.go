package main

import "fmt"

var hashMap = make(map[string]int)

func main() {
	addValuesToMap()
	fmt.Println(getValueFromMap("valueOne"))
	fmt.Println(getValueFromMap("valueTwo"))
}

func addValuesToMap() {
	hashMap["valueOne"] = 1
	hashMap["valueTwo"] = 2
}

func getValueFromMap(key string) (result int) {
	return hashMap[key]
}
