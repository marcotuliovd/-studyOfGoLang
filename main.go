package main

import (
	"pix/cmd"
	donutfan "pix/donutFan"
)

func ReturnMeu(amount int) int {
	r := cmd.Index(amount)
	return r
}

func main() {
	ReturnMeu(10)
	donutfan.TomaCafeComDonut()
}