package main

import (
	"fmt"
	"github.com/ghaneakbar4/GolangConvertPack/convert"
)


func main() {
	res :=convert.PrsToInt("2")
    res1 := convert.ToString(2)
	fmt.Print(res , res1)
}