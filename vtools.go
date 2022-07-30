// Package vtools created by guidoenr, is a set of tools to code faster. Specifically written for veritone-engines/*
package vtools

import (
	"fmt"
	"reflect"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// Print the input in a pretty way
func Print(name string, in interface{}) {
	fmt.Println("-----------------------------")
	fmt.Printf(Cyan+"%s: %v"+Reset+" \n", name, in)
	fmt.Printf(Gray+"type: %v"+Reset+"\n", reflect.TypeOf(in))
	fmt.Println("-----------------------------")
}
