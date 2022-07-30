// Package vtools created by guidoenr, is a set of tools to code faster. Specifically written for veritone-engines/*
package vtools

import (
	"github.com/fatih/color"
	"reflect"
)

// Print the input in a pretty way
func Print(name string, in interface{}) {
	color.Cyan("------------")
	color.Cyan(name+": ", in)
	color.Cyan("type: ", reflect.TypeOf(in))
	color.Cyan("------------")
}
