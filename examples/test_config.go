package main

import (
	"fmt"

	"github.com/5gif/config"
	"github.com/wiless/vlib"
)

func main() {

	config.SetDir("../json", "results")

	fmt.Println("\n\n=========== Use method 1 ===========")
	var nrcfg config.NRconfig                // loads the 3GPP related parameters for the RURAL Evaluation Config. A (of ITU-R/WP5D)
	path1 := "../json/3GPP_RMa_configA.json" // Add absolute path to the json file
	vlib.LoadStructure(path1, &nrcfg)
	fmt.Printf("\nNR Config = %#v", nrcfg)

	// var ituconfig config.ITUconfig          // loads the 3GPP related parameters for the RURAL Evaluation Config. A (of ITU-R/WP5D)
	// path2 := "../json/ITU_RMa_configA.json" // Add absolute path to the json file
	// vlib.LoadStructure(path2, &ituconfig)
	// fmt.Printf("\n\nITU-R Config = %#v", ituconfig)

	fmt.Println("\n\n=========== Use method 2 ===========")
	var icfg config.ITUconfig
	icfg.Read("ITU_RMa_configB.json")
	fmt.Printf("\n\nITU-R Method 2 Config = %#v\n", icfg)
	icfg.Save()

	fmt.Println("\n\n=========== Use method 3 ===========")
	icfg2 := config.ReadITUConfig("ITU_RMa_configA.json", "../json")
	fmt.Printf("\n\nITU-R Method 2 Config = %#v\n", icfg2)
	icfg2.Save()

}
