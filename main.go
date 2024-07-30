package main

import (
	"fmt"
	"go-hello/greeting"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.HiYellowString(greeting.Get()))
}
