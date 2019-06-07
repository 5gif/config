package main

import "fmt"
import "github.com/5gif/config/datatypes"

func main() {
	var confg datatypes.Params3GPP
	confg.FcGHz = 10
	fmt.Println(confg.FcGHz)
}
