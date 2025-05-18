package main

import (
	"fmt"
	"learn/mypackage"
	"learn/mypackage/crud"
	"mymodule"
)

func main() {
	fmt.Println("MAIN TEST")
	mypackage.SayPackage()
	crud.SaySubPackage()

	mymodule.SayMyModule()
}
