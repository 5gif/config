package main

import (
	"fmt"
	"os"

	"github.com/5gif/config"
	"github.com/wiless/vlib"
)

func main() {
	pwd, _ := os.Getwd()
	cfgDir := pwd + "/" + "../json"
	fmt.Println(cfgDir)
	config.SetDir(cfgDir, "results")

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
	icfg2, err2 := config.ReadITUConfig("ITU_RMa_configA.json")
	_ = err2
	fmt.Printf("\n\nITU-R Method 3 Config = %#v\n", icfg2)
	icfg2.Save()

	fmt.Println("\n\n=========== Use method 3 ===========")
	icfg3, err3 := config.ReadSIMConfig("SIM_RMa_configA.json")
	_ = err3
	fmt.Printf("\n\nSIM-R Method 3 Config = %#v\n", icfg3)
	icfg3.Save()

	fmt.Println("\n\n=========== Use method 3 ===========")
	icfg4, err4 := config.ReadNRConfig("3GPP_UMa_configA.json")
	_ = err4
	fmt.Printf("\n\n3GPP-U Method 3 Config = %#v\n", icfg4)
	icfg4.Save()

}
