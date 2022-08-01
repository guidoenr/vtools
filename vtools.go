// Package vtools created by guidoenr, is a set of tools to code faster. Specifically written for veritone-engines /*
package vtools

import (
	"fmt"
	"reflect"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var purple = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

// Print the input in a pretty way, and prints every type of var
func Print(name string, in interface{}) {
	fmt.Println("-----------------------------")
	fmt.Printf(cyan+"%s: %v"+reset+" \n", name, in)
	fmt.Printf(gray+"type: %v"+reset+"\n", reflect.TypeOf(in))
	fmt.Println("-----------------------------")
}
