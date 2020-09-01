package main

import custompackage "./MyCustomPackages"

func main() {
	var structType1 = custompackage.CustomTypeOne{}
	var structType2 = custompackage.CustomTypeTwo{}
	var structType3 = custompackage.CustomTypeThree{}

	structType1.CustomMethodOne()
	structType2.CustomMethodTwo()
	structType3.CustomMethodThree()
}
