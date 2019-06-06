package main

import "github.com/zchee/color"

func main() {
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Hello Gophers!")
}
