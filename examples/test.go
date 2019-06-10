package main

import (
	"fmt"

	"github.com/5gif/config"

	"github.com/wiless/vlib"
)

func main() {
	var confg config.NRconfig
	var path2file = "/home/krishnan2098/goapps/src/github.com/5gif/config/json/3GPP_Ru_configA.json" //Add absolute path to the json file
	vlib.LoadStructure(path2file, &confg)
	fmt.Println(confg.FcGHz)
}
