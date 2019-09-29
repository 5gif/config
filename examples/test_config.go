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
	config.PrintStructsPretty(nrcfg)

	fmt.Println("\n\n=========== Use method 2 ===========")
	var icfg config.ITUconfig
	icfg.Read("ITU_RMa_configB.json")
	config.PrintStructsPretty(icfg)
	icfg.Save()

	fmt.Println("\n\n=========== Use method 3 ===========")
	icfg2, err2 := config.ReadITUConfig("ITU_RMa_configA.json")
	_ = err2
	config.PrintStructsPretty(icfg2)
	icfg2.Save()

	fmt.Println("\n\n=========== Use method 4 ===========")
	icfg3, err3 := config.ReadSIMConfig("SIM_RMa_configA.json")
	_ = err3
	config.PrintStructsPretty(icfg3)
	icfg3.Save()

	fmt.Println("\n\n=========== Use method 5 ===========")
	icfg4, err4 := config.ReadNRConfig("3GPP_UMa_configA.json")
	_ = err4
	config.PrintStructsPretty(icfg4)
	icfg4.Save()

	fmt.Println("\n\n=========== Use method 6 ===========")
	// The flow of main() to read calls a setup function using an input json file
	var err error
	var input = "default.json"
	var appsetting config.AppSetting
	appsetting.FromJSON(input)
	app, err := appsetting.LoadApp()

	// i, n, s, _, err = config.DefaultApp.Load(input)
	if err == nil {
		config.PrintStructsPretty(app.ITUcfg)
		config.PrintStructsPretty(app.NRcfg)
		config.PrintStructsPretty(app.SIMcfg)
	}
}
