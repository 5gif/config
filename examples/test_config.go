package main

import (
	"fmt"

	"github.com/5gif/config"

	"github.com/wiless/vlib"
)

func main() {

	var nrcfg config.NRconfig                // loads the 3GPP related parameters for the RURAL Evaluation Config. A (of ITU-R/WP5D)
	path1 := "../json/3GPP_RMa_configA.json" // Add absolute path to the json file
	vlib.LoadStructure(path1, &nrcfg)
	fmt.Printf("\nNR Config = %#v", nrcfg)

	var ituconfig config.ITUconfig          // loads the 3GPP related parameters for the RURAL Evaluation Config. A (of ITU-R/WP5D)
	path2 := "../json/ITU_RMa_configA.json" // Add absolute path to the json file
	vlib.LoadStructure(path2, &ituconfig)
	fmt.Printf("\n\nITU-R  Config = %#v", ituconfig)
}
